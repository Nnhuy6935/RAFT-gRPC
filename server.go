package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"main/services/raft"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const (
	// HEART_BEAT_TIMEOUT = 1000 // đơn vị tính băng ms
	// ELECTION_TIMEOUT   = 3000
	HEART_BEAT_TIMEOUT = 1000 // đơn vị tính băng ms
	ELECTION_TIMEOUT   = 3000
	CANDIDATE          = 1
	FOLLOWER           = 0
	LEADER             = 2
)

var restartChan chan bool

type Server struct {
	raft.UnimplementedRaftServer
	id               int32
	leaderId         int32
	term             int32
	votedFor         int32
	log              []string
	mu               sync.Mutex
	nodes            []raft.NodeInfo
	consulClient     *api.Client
	heartbeatTimeout int
	electionTimeout  int
	status           int
	value            int32
	temp             int32
}

func NewServer(id int32) *Server {
	return &Server{
		id:               id,
		term:             0,
		votedFor:         -1,
		log:              []string{},
		status:           FOLLOWER,
		heartbeatTimeout: HEART_BEAT_TIMEOUT,
		electionTimeout:  ELECTION_TIMEOUT,
		value:            0,
		temp:             0,
	}
}

func (s *Server) RequestVote(ctx context.Context, req *raft.RequestVoteRequest) (*raft.RequestVoteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// log.Printf("TERM %d : GET REQUEST VOTE FOR %d", req.Term, req.CandidateId)
	s.log = append(s.log, fmt.Sprintf("%s: TERM %d : GET REQUEST VOTE FOR %d", time.Now().Format("2006/01/02 15:04:05"), req.Term, req.CandidateId))
	if req.Term > s.term {
		s.term = req.Term
		s.votedFor = -1 // Reset votedFor
	}

	if s.leaderId != -1 || s.status == LEADER || s.status == CANDIDATE {
		s.electionTimeout = ELECTION_TIMEOUT
		restartRoroutine(restartChan)
		return &raft.RequestVoteResponse{Term: s.term, VoteGranted: false}, nil
	}

	if (s.votedFor == -1 || s.votedFor == req.CandidateId) && req.Term == s.term {
		s.votedFor = req.CandidateId
		s.electionTimeout = ELECTION_TIMEOUT
		restartRoroutine(restartChan)
		return &raft.RequestVoteResponse{Term: s.term, VoteGranted: true}, nil
	}

	s.electionTimeout = ELECTION_TIMEOUT
	restartRoroutine(restartChan)
	return &raft.RequestVoteResponse{Term: s.term, VoteGranted: false}, nil
}

func (s *Server) AppendEntries(ctx context.Context, req *raft.AppendEntriesRequest) (*raft.AppendEntriesResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Term < s.term {
		return &raft.AppendEntriesResponse{Term: s.term, Success: false}, nil
	}

	s.term = req.Term
	// Xử lý log entries tại đây
	return &raft.AppendEntriesResponse{Term: s.term, Success: true}, nil
}

// Đăng ký node
func (s *Server) RegisterNode(ctx context.Context, node *raft.NodeInfo) (*raft.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	//register node with consul

	// Tách địa chỉ IP và cổng
	parts := strings.Split(node.Address, ":")
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Printf("TERM %d: error when convert from int to string %v", s.term, err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%d", node.Id), // Sử dụng địa chỉ làm ID
		Name:    "agent_service",
		Port:    port,
		Address: parts[0],
	}

	error := s.consulClient.Agent().ServiceRegister(registration)
	if error != nil {
		log.Printf("TERM %d: error in ServiceRegister: %v", s.term, error)
		s.log = append(s.log, fmt.Sprintf("%s: TERM %d: error in ServiceRegister: %v", time.Now().Format("2006/01/02 15:04:05"), s.term, error))
		return &raft.Response{Success: false, Message: error.Error()}, nil
	}

	s.nodes = append(s.nodes, *node) // Thêm node vào danh sách
	s.log = append(s.log, fmt.Sprintf("%s: TERM %d: ServiceRegister success!", time.Now().Format("2006/01/02 15:04:05"), s.term))
	return &raft.Response{Success: true, Message: "Node registered successfully"}, nil
}

// Lấy danh sách node trong mạng
func (s *Server) GetNodeList(ctx context.Context, empty *raft.Empty) (*raft.NodeList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	//lấy danh sách các node từ consul
	services, err := s.consulClient.Agent().Services()
	if err != nil {
		return nil, err
	}

	var nodes []raft.NodeInfo
	for _, service := range services {
		// get service id and convert to int32
		id, err := strconv.Atoi(service.ID)
		if err != nil {
			log.Printf("TERM %d: get list node return error %v", s.term, err)
		}
		nodes = append(nodes, raft.NodeInfo{Id: int32(id), Address: fmt.Sprintf("%s:%d", service.Address, service.Port)}) // ID có thể được điều chỉnh
	}
	//CONVERT
	var nodePointers []*raft.NodeInfo

	for i := range nodes {
		nodePointers = append(nodePointers, &nodes[i])
	}
	//END CONVERT
	return &raft.NodeList{Nodes: nodePointers}, nil
}
func (s *Server) RequestLeader(ctx context.Context, req *raft.RequestLeaderRequest) (*raft.RequestLeaderResponse, error) {
	// log.Printf("TERM %d: RECEIVED REQUEST LEADER from %d -- CURRENT LEADER IN SERVER IS %d - STATUS %d - SERVER TERM %d", req.Term, req.LeaderId, s.leaderId, s.status, s.term)
	// kiểm tra leaderid
	if s.status != LEADER && (s.term == req.Term || s.leaderId == -1) {
		s.status = FOLLOWER
		s.electionTimeout = ELECTION_TIMEOUT
		s.heartbeatTimeout = HEART_BEAT_TIMEOUT
		s.term = req.Term
		s.leaderId = req.LeaderId
		if s.value != req.Value {
			log.Printf("TERM %d: Update value", s.term, int(req.Value))
		}
		s.value = req.Value
		// log.Printf("TERM %d: AFTER RECEIVED REQUEST LEADER from %d -- CURRENT LEADER IN SERVER IS %d - STATUS %d", req.Term, req.LeaderId, s.leaderId, s.status)
		s.log = append(s.log, fmt.Sprintf("%s : TERM %d: RECEIVED REQUEST LEADER from %d -- CURRENT LEADER IN SERVER IS %d - STATUS %d - SERVER TERM %d", time.Now().Format("2006/01/02 15:04:05"), req.Term, req.LeaderId, s.leaderId, s.status, s.term))

		restartRoroutine(restartChan)
		return &raft.RequestLeaderResponse{
			Term:     s.term,
			LeaderId: s.leaderId,
		}, nil
	}
	return nil, nil
}

func (s *Server) startElection() {
	log.Printf("TERM %d: START ELECTION TO BECOME LEADER", s.term)
	s.log = append(s.log, fmt.Sprintf("%s: TERM %d: START ELECTION TO BECOME LEADER", time.Now().Format("2006/01/02 15:04:05"), s.term))
	list, err := s.GetNodeList(context.Background(), &raft.Empty{})
	if err != nil {
		log.Println(err)
		return
	}
	countVote := 0
	countResponse := 0
	s.term += 1
	for i := 0; i < len(list.Nodes); i++ {
		// bỏ qua và không gửi request tới chính nó
		if s.id == list.Nodes[i].Id {
			countVote++
			continue
		}

		conn, err := grpc.Dial(list.Nodes[i].Address, grpc.WithInsecure())
		if err != nil {
			log.Printf("TERM %d: Failed to connect to node %s: %v", s.term, list.Nodes[i].Address, err)
			s.log = append(s.log, fmt.Sprintf("%s: TERM %d: Failed to connect to node %s: %v", time.Now().Format("2006/01/02 15:04:05"), s.term, list.Nodes[i].Address, err))
			continue
		}
		// đảm bảo đóng kết nối ngay sau khi sử dụng
		defer conn.Close()

		client := raft.NewRaftClient(conn)

		req := &raft.RequestVoteRequest{
			Term:        s.term,
			CandidateId: s.id,
		}

		resp, err := client.RequestVote(context.Background(), req)
		if err != nil {
			s.log = append(s.log, fmt.Sprintf("%s: TERM %d: Error while requesting vote from node %s: %v", time.Now().Format("2006/01/02 15:04:05"), s.term, list.Nodes[i].Address, err))
			log.Printf("TERM %d: Error while requesting vote from node %s: %v", s.term, list.Nodes[i].Address, err)
			continue
		}
		if resp.Term == s.term {
			countResponse++
		}

		if resp.VoteGranted {
			log.Printf("TERM %d: Node %s granted vote to node %d", s.term, list.Nodes[i].Address, s.leaderId)
			s.log = append(s.log, fmt.Sprintf("%s: TERM %d: Node %s granted vote to node %d", time.Now().Format("2006/01/02 15:04:05"), s.term, list.Nodes[i].Address, s.leaderId))
			countVote++
		} else {
			log.Printf("TERM %d: Node %s denied vote to node %d", s.term, list.Nodes[i].Address, s.leaderId)
			s.log = append(s.log, fmt.Sprintf("%s: TERM %d: Node %s denied vote to node %d", time.Now().Format("2006/01/02 15:04:05"), s.term, list.Nodes[i].Address, s.leaderId))
		}
	}
	// if countVote > len(list.Nodes)/2 {
	if countVote > countResponse/2 {
		s.status = LEADER
		s.heartbeatTimeout = HEART_BEAT_TIMEOUT
		s.leaderId = s.id
	} else {
		s.status = FOLLOWER
		s.electionTimeout = ELECTION_TIMEOUT
		s.leaderId = -1
	}
}

// hàm thực hiện việc gửi tín hiệu định kỳ tới các node follower
func (s *Server) sendHearbeatMessage() {
	log.Printf("TERM %d: SEND HEART BEAT MESSAGE TO FOLLOWER", s.term)
	s.log = append(s.log, fmt.Sprintf("%s: ERM %d: SEND HEART BEAT MESSAGE TO FOLLOWER", time.Now().Format("2006/01/02 15:04:05"), s.term))
	// duyệt qua từng node trong mạng để request leader mới
	list, err := s.GetNodeList(context.Background(), &raft.Empty{})
	if err != nil {
		log.Println(err)
		return
	}

	// gửi requestLeader tới từng node
	for i := 0; i < len(list.Nodes); i++ {
		//bỏ qua việc gửi request tới chính nó
		if s.id == list.Nodes[i].Id {
			continue
		}

		conn, err := grpc.Dial(list.Nodes[i].Address, grpc.WithInsecure())
		if err != nil {
			log.Printf("TERM %d: Failed to connect to node %s: %v", s.term, list.Nodes[i].Address, err)
			continue
		}
		// đảm bảo đóng kết nối ngay sau khi sử dụng
		defer conn.Close()

		client := raft.NewRaftClient(conn)

		//setup requestLeaderRequest
		req := &raft.RequestLeaderRequest{
			Term:     s.term,
			LeaderId: s.id,
			Value:    s.value,
		}
		response, err := client.RequestLeader(context.Background(), req)
		if err != nil {
			log.Printf("TERM %d: Error while requesting leader from node %s", s.term, list.Nodes[i].Address)
			continue
		}
		log.Printf("TERM %d: RESPONSE %s", response.Term, response.String())
	}
	s.heartbeatTimeout = HEART_BEAT_TIMEOUT
}

// node client gửi request tới leader để update value của network
func (s *Server) sendRequestToLeader(data int32) {
	list, err := s.GetNodeList(context.Background(), &raft.Empty{})
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < len(list.Nodes); i++ {
		if list.Nodes[i].Id != s.leaderId {
			continue
		}
		//kết nối với leader node để gửi dữ liệu
		conn, err := grpc.Dial(list.Nodes[i].Address, grpc.WithInsecure())
		if err != nil {
			log.Printf("TERM %d: Failed to connect to node %s: %v", s.term, list.Nodes[i].Address, err)
			continue
		}
		defer conn.Close()

		client := raft.NewRaftClient(conn)

		req := &raft.RequestValueToLeaderRequest{
			NodeRequestId: s.id,
			RequestValue:  data,
			Term:          s.term,
		}

		response, err := client.RequestValueToLeader(context.Background(), req)
		if err != nil {
			log.Printf("TERM %d: Error while node %d requesting leader to change value", s.term, list.Nodes[i].Address)
			continue
		}
		log.Printf("TERM %d: RESPONSE %s", s.term, response.Message)
	}
}

func (s *Server) collectVoteChangeValue() (*raft.Response, error) {
	if s.temp != s.value {
		list, err := s.GetNodeList(context.Background(), &raft.Empty{})
		if err != nil {
			log.Println(err)
			return &raft.Response{
				Success: false,
				Message: "failed to get list node in network",
			}, nil
		}

		totalVote := 0
		grantedVote := 0

		for i := 0; i < len(list.Nodes); i++ {
			if list.Nodes[i].Id == s.leaderId {
				continue
			}

			conn, err := grpc.Dial(list.Nodes[i].Address, grpc.WithInsecure())
			if err != nil {
				log.Printf("TERM %d: Failed to connect to node %s: %v", s.term, list.Nodes[i].Address, err)
				continue
			}
			// đảm bảo đóng kết nối ngay sau khi sử dụng
			defer conn.Close()

			client := raft.NewRaftClient(conn)

			// s.term++
			req := &raft.RequestVoteChangeValueRequest{
				LeaderId:     s.id,
				Term:         s.term,
				RequestValue: s.temp,
			}

			response, err := client.RequestVoteChangeValue(context.Background(), req)
			if err != nil {
				log.Printf("TERM %d : Failed to get vote from client %s", s.term, list.Nodes[i].Address)
				continue
			}

			totalVote++
			if response.VoteGranted == true {
				grantedVote++
			}

		}
		if grantedVote > totalVote/2 {
			log.Printf("TERM %d: CHANGE ACCEPT", s.term)
			s.value = s.temp
			return &raft.Response{
				Success: true,
				Message: "Change accept",
			}, nil
		}

		log.Printf("TERM %d: CHANGE REFUSE with granted %d per total %d", s.term, grantedVote, totalVote)
		return &raft.Response{
			Success: false,
			Message: "Change Refuse",
		}, nil

	}
	return &raft.Response{
		Success: false,
		Message: "request value equal with current value in network",
	}, nil
}

// leader xử lý request từ client khi nhận được request
func (s *Server) RequestValueToLeader(ctx context.Context, req *raft.RequestValueToLeaderRequest) (*raft.Response, error) {
	log.Printf("TERM %d: RECEIVED CHANGE VALUE FROM %d with term %d and value is %d", s.term, req.NodeRequestId, req.Term, req.RequestValue)
	if req.Term == s.term {
		s.temp = req.RequestValue
		return s.collectVoteChangeValue()
	}

	return &raft.Response{
		Success: false,
		Message: fmt.Sprintf("TERM %d: Term is not match with current term %d", s.term, req.Term),
	}, nil

}

func (s *Server) RequestVoteChangeValue(ctx context.Context, req *raft.RequestVoteChangeValueRequest) (*raft.RequestVoteChangeValueResponse, error) {

	log.Printf("TERM %d: GET request vote change value from leader with request term %d and value %d", s.term, req.Term, req.RequestValue)
	if req.Term > s.term {
		s.term = req.Term
	}
	if req.LeaderId == s.leaderId && req.RequestValue != s.value {
		return &raft.RequestVoteChangeValueResponse{
			LeaderId:    s.leaderId,
			Term:        s.term,
			VoteGranted: true,
		}, nil
	}

	return &raft.RequestVoteChangeValueResponse{
		LeaderId:    s.leaderId,
		Term:        s.term,
		VoteGranted: false,
	}, nil
}

type Service struct {
	consulClient *api.Client
}

func NewService(port int) (*Service, error) {
	// Tạo cấu hình cho Consul
	config := api.DefaultConfig()

	// chỉ định địa chỉ của consul
	config.Address = "127.0.0.1:8500"

	// Khởi tạo Consul client
	client, err := api.NewClient(config)
	if err != nil {
		log.Println("Unable to contact Service Discovery.")
		return nil, err
	}

	return &Service{consulClient: client}, nil
}

func restartRoroutine(restart chan bool) {
	go func() {
		restart <- true
	}()
}

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("STARTING... : Usage: %s <node_id> <port>", os.Args[0])
	}

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("STARTING... : Invalid node_id: %v", err)
	}

	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("STARTING... : Invalid port: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf("STARTING... : failed to listen: %v", err)
	} else {
		log.Println(lis.Addr())
	}
	service, err := NewService(port)
	if err != nil {
		log.Fatalf("STARTING... : Error creating service: %v", err)
	}

	srv := NewServer(int32(id))
	srv.consulClient = service.consulClient
	grpcServer := grpc.NewServer()
	raft.RegisterRaftServer(grpcServer, srv)

	//đảm bảo khởi tạo trước khi sử dụng
	srv.nodes = make([]raft.NodeInfo, 0)

	node := &raft.NodeInfo{Id: int32(id), Address: fmt.Sprintf("127.0.0.1:%d", port)}
	response, err := srv.RegisterNode(context.Background(), node)
	if err != nil {
		log.Printf("TERM %d: Error when register node in server", srv.term)
		srv.log = append(srv.log, fmt.Sprintf("%s : TERM %d: Error when register node in server", time.Now().Format("2006/01/02 15:04:05"), srv.term))
	}

	log.Println(response.Message)

	restartChan = make(chan bool)

	// lắng nghe các request từ các node khác
	go func() {
		for {
			err = grpcServer.Serve(lis)
			if err != nil {
				log.Fatalf("TERM %d: Failed to serve: %v", srv.term, err)
				srv.log = append(srv.log, fmt.Sprintf("%s : TERM %d: Failed to serve: %v", time.Now().Format("2006/01/02 15:04:05"), srv.term, err))
			}
		}
	}()
	go func(restart chan bool) {
		for {
			select {
			case <-restart:
				{
					break
				}
			default:
				{
					switch srv.status {
					case FOLLOWER:
						{
							// log.Printf("TERM %d: status FOLLOWER", srv.term)
							// trường hợp chưa chọn leader thì sẽ thực hiện đếm ngược thời gian để trở thành candidate
							if srv.leaderId == -1 {
								// Chờ một khoảng thời gian electionTimeout
								startTime := time.Now()
								for srv.electionTimeout > 0 {
									if srv.leaderId != -1 {
										break
									}
									currentTime := time.Now()
									duration := currentTime.Sub(startTime).Milliseconds()
									srv.electionTimeout -= int(duration)
									startTime = currentTime
								}
								if srv.electionTimeout <= 0 {
									srv.status = CANDIDATE
								}
								break
							}
							//trường hợp đã chọn leader thì sẽ thực hiện đếm ngược thời gian để trở thành follower
							startTime := time.Now()
							for srv.heartbeatTimeout > 0 {
								currentTime := time.Now()
								duration := currentTime.Sub(startTime).Milliseconds()
								srv.heartbeatTimeout -= int(duration)
								startTime = currentTime
							}
							if srv.heartbeatTimeout <= 0 {
								srv.electionTimeout = ELECTION_TIMEOUT
								srv.leaderId = -1
							}
							break
						}
					case CANDIDATE:
						{
							// log.Printf("TERM %d: status CANDIATE", srv.term)
							// bắt đầu gửi request bầu cử
							srv.startElection() // Gọi runElectionProcess từ instance của server
							break
						}
					case LEADER:
						{
							log.Printf("TERM %d: status LEADER", srv.term)
							startTime := time.Now()
							for srv.heartbeatTimeout > 0 {
								currentTime := time.Now()
								duration := currentTime.Sub(startTime).Milliseconds()
								srv.heartbeatTimeout -= int(duration)
								startTime = currentTime
							}
							if srv.heartbeatTimeout <= 0 {
								srv.sendHearbeatMessage()
							}
							break
						}
					}
					break
				}
			}
		}

	}(restartChan)

	// lắng nghe người dùng từ console
	go func() {
		for {
			if srv.status != LEADER {
				reader := bufio.NewReader(os.Stdin)
				text, _ := reader.ReadString('\n')
				text = strings.TrimSpace(text) // Xóa ký tự xuống dòng
				changeVal, err := strconv.Atoi(text)
				if err != nil {
					log.Printf("TERM %d: Failed to convert string to int %v", srv.term, err)
				}

				log.Printf("REQUEST CHANGE VALUE TO %d", changeVal)
				srv.sendRequestToLeader(int32(changeVal))
			}
		}
	}()

	// Để giữ cho main goroutine không kết thúc
	select {}

}
