package main

import (
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
	HEART_BEAT_TIMEOUT = 10000 // đơn vị tính băng ms
	ELECTION_TIMEOUT   = 30000
	CANDIDATE          = 1
	FOLLOWER           = 0
	LEADER             = 2
)

type Server struct {
	raft.UnimplementedRaftServer
	id               int32
	leaderId         int32
	term             int32
	votedFor         int32
	log              []string
	peers            []string
	mu               sync.Mutex
	nodes            []raft.NodeInfo
	consulClient     *api.Client
	heartbeatTimeout int
	electionTimeout  int
	status           int
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
	}
}

func (s *Server) RequestVote(ctx context.Context, req *raft.RequestVoteRequest) (*raft.RequestVoteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// log.Printf("request vote with information {term: %d , candiate %s} with server information {voteFor: %d, term %d}", req.Term, req.CandidateId, s.votedFor, s.term)

	if s.leaderId != -1 {
		s.electionTimeout = ELECTION_TIMEOUT
		return &raft.RequestVoteResponse{Term: s.term, VoteGranted: false}, nil
	}
	if req.Term > s.term {
		s.term = req.Term
		s.votedFor = -1 // Reset votedFor
	}
	if (s.votedFor == -1 || s.votedFor == req.CandidateId) && req.Term == s.term {
		s.votedFor = req.CandidateId
		s.electionTimeout = ELECTION_TIMEOUT
		return &raft.RequestVoteResponse{Term: s.term, VoteGranted: true}, nil
	}

	s.electionTimeout = ELECTION_TIMEOUT
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

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", 8000+s.id))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	raft.RegisterRaftServer(grpcServer, s)

	log.Printf("Server %d is listening on port %d", s.id, 8000+s.id)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func (s *Server) startElection() {

	list, err := s.GetNodeList(context.Background(), &raft.Empty{})
	if err != nil {
		log.Println(err)
		return
	}
	countVote := 0
	for i := 0; i < len(list.Nodes); i++ {
		// bỏ qua và không gửi request tới chính nó
		if s.id == list.Nodes[i].Id {
			countVote++
			continue
		}

		conn, err := grpc.Dial(list.Nodes[i].Address, grpc.WithInsecure())
		if err != nil {
			log.Printf("Failed to connect to node %s: %v", list.Nodes[i].Address, err)
			continue
		}
		// đảm bảo đóng kết nối ngay sau khi sử dụng
		defer conn.Close()

		client := raft.NewRaftClient(conn)
		// log.Printf("REQUEST VOTE INFORMATION term(%d) and candidate id(%d)", s.term, int(s.leaderId))
		req := &raft.RequestVoteRequest{
			Term:        s.term + 1,
			CandidateId: s.leaderId,
		}

		// log.Printf("start request vote for client ")
		resp, err := client.RequestVote(context.Background(), req)
		if err != nil {
			log.Printf("Error while requesting vote from node %s: %v", list.Nodes[i].Address, err)
			continue
		}

		if resp.VoteGranted {
			log.Printf("Node %s granted vote to node %d", list.Nodes[i].Address, s.leaderId)
			countVote++
		} else {
			log.Printf("Node %s denied vote to node %d", list.Nodes[i].Address, s.leaderId)
		}
	}
	if countVote > len(list.Nodes)/2 {
		log.Printf("Node %d has received enough votes to become the leader", s.id)
		s.status = LEADER
		s.heartbeatTimeout = HEART_BEAT_TIMEOUT
	} else {
		log.Printf("Node %d did not receive enough votes to become the leader", s.id)
		s.status = FOLLOWER
		s.electionTimeout = ELECTION_TIMEOUT
	}
}

// hàm thực hiện việc gửi tín hiệu định kỳ tới các node follower
func (s *Server) sendHearbeatMessage() {
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
			log.Printf("Failed to connect to node %s: %v", list.Nodes[i].Address, err)
			continue
		}
		// đảm bảo đóng kết nối ngay sau khi sử dụng
		defer conn.Close()

		client := raft.NewRaftClient(conn)

		//setup requestLeaderRequest
		req := &raft.RequestLeaderRequest{
			Term: s.term,
		}
		response, err := client.RequestLeader(context.Background(), req)
		if err != nil {
			log.Printf("Error while requesting leader from node %s", list.Nodes[i].Address)
		}
		log.Printf(response.String())
	}
	s.heartbeatTimeout = HEART_BEAT_TIMEOUT
}

func (s *Server) RequestLeader(ctx context.Context, req *raft.RequestLeaderRequest) (*raft.RequestLeaderResponse, error) {
	// Logic để chọn leader
	// Ở đây, bạn có thể thực hiện logic của riêng mình để xác định leader

	// kiểm tra leaderid
	if req.Term == s.term && s.leaderId != -1 {
		s.status = FOLLOWER
		s.heartbeatTimeout = HEART_BEAT_TIMEOUT
		s.term = req.Term
		s.leaderId = req.LeaderId
		return &raft.RequestLeaderResponse{
			Term:     s.term,
			LeaderId: s.leaderId,
		}, nil
	}
	return nil, nil

}

func startNode(id int32, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	raft.RegisterRaftServer(s, &Server{leaderId: id, term: 1}) // Khởi tạo leaderId và term
	log.Printf("Node %d started on port %d", id, port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) runElectionProcess() {
	// Logic để xác định khi nào cần bắt đầu bầu cử
	// Ví dụ: sau một khoảng thời gian nhất định hoặc khi không nhận được heartbeat từ leader
	s.startElection()
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
		log.Printf("error when convert from int to string %v", err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%d", node.Id), // Sử dụng địa chỉ làm ID
		Name:    "agent_service",
		Port:    port,
		Address: parts[0],
	}

	error := s.consulClient.Agent().ServiceRegister(registration)
	if error != nil {
		log.Printf("error in ServiceRegister: %v", error)
		return &raft.Response{Success: false, Message: error.Error()}, nil
	}

	s.nodes = append(s.nodes, *node) // Thêm node vào danh sách
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
			log.Printf("error %v", err)
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

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <node_id> <port>", os.Args[0])
	}

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid node_id: %v", err)
	}

	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println(lis.Addr())
	}
	service, err := NewService(port)
	if err != nil {
		log.Fatalf("Error creating service: %v", err)
	}

	srv := NewServer(int32(id))
	srv.consulClient = service.consulClient
	grpcServer := grpc.NewServer()
	raft.RegisterRaftServer(grpcServer, srv)

	// lắng nghe các request từ các node khác
	go func() {
		for {
			// log.Println("Starting gRPC server...")
			err = grpcServer.Serve(lis)
			if err != nil {
				log.Fatalf("Failed to serve: %v", err)
			}
			// log.Println("gRPC server started successfully.")
		}
	}()

	//đảm bảo khởi tạo trước khi sử dụng
	srv.nodes = make([]raft.NodeInfo, 0)

	node := &raft.NodeInfo{Id: int32(id), Address: fmt.Sprintf("127.0.0.1:%d", port)}
	response, err := srv.RegisterNode(context.Background(), node)
	if err != nil {
		log.Println(err)
	}

	log.Println(response.Message)
	go func() {
		for {
			switch srv.status {
			case FOLLOWER:
				{
					log.Printf("status FOLLOWER")
					// trường hợp chưa chọn leader thì sẽ thực hiện đếm ngược thời gian để trở thành candidate
					if srv.leaderId == -1 {
						// Chờ một khoảng thời gian electionTimeout
						startTime := time.Now()
						for srv.electionTimeout > 0 {
							currentTime := time.Now()
							duration := currentTime.Sub(startTime).Milliseconds()
							srv.electionTimeout -= int(duration)
							startTime = currentTime
						}
						if srv.electionTimeout <= 0 {
							// bắt đầu gửi request bầu cử
							log.Printf("START ELECTION REQUEST")
							srv.runElectionProcess() // Gọi runElectionProcess từ instance của server
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
					// log.Printf("status CANDIATE")
					break
				}
			case LEADER:
				{
					log.Printf("status LEADER")
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
		}

	}()

	// Để giữ cho main goroutine không kết thúc
	select {}

}
