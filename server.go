package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	// "time"

	"main/yourmodule/raft" // Thay thế bằng module của bạn

	"google.golang.org/grpc"
)

// ----------- CENTRAL NODE ----------------

// ------------------ SERVER ----------------

type Server struct {
	raft.UnimplementedRaftServer
	id       int32
	leaderId int32
	term     int32
	votedFor int32
	log      []string
	peers    []string
	mu       sync.Mutex
}

func NewServer(id int32, peers []string) *Server {
	return &Server{
		id:       id,
		term:     0,
		votedFor: -1,
		log:      []string{},
		peers:    peers,
	}
}

func (s *Server) RequestVote(ctx context.Context, req *raft.RequestVoteRequest) (*raft.RequestVoteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Term > s.term {
		s.term = req.Term
		s.votedFor = -1 // Reset votedFor
	}

	if (s.votedFor == -1 || s.votedFor == req.CandidateId) && req.Term == s.term {
		s.votedFor = req.CandidateId
		return &raft.RequestVoteResponse{Term: s.term, VoteGranted: true}, nil
	}

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
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000+s.id))
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
	for index := 0; index < len(s.peers); index++ {
		log.Printf(s.peers[index])
	}
	log.Printf("Number of peers in network %d", len(s.peers))
	for _, peer := range []int32{2, 3} { // Giả định rằng bạn có 3 node với ID 1, 2, 3
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", 50050+peer), grpc.WithInsecure())
		if err != nil {
			log.Printf("Failed to connect to node %d: %v", peer, err)
			continue
		}
		defer conn.Close()

		client := raft.NewRaftClient(conn)
		req := &raft.RequestVoteRequest{
			Term:        s.term,
			CandidateId: s.leaderId,
		}

		resp, err := client.RequestVote(context.Background(), req)
		if err != nil {
			log.Printf("Error while requesting vote from node %d: %v", peer, err)
			continue
		}

		if resp.VoteGranted {
			log.Printf("Node %d granted vote to node %d", peer, s.leaderId)
		} else {
			log.Printf("Node %d denied vote to node %d", peer, s.leaderId)
		}
	}
}
func (s *Server) RequestLeader(ctx context.Context, req *raft.RequestLeaderRequest) (*raft.RequestLeaderResponse, error) {
	// Logic để chọn leader
	// Ở đây, bạn có thể thực hiện logic của riêng mình để xác định leader
	return &raft.RequestLeaderResponse{
		Term:     s.term,
		LeaderId: s.leaderId,
	}, nil
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

func main() {
	// peers := []string{"localhost:8001", "localhost:8002", "localhost:8003"} // Thay đổi theo số lượng node
	// for i := 0; i < len(peers); i++ {
	// 	server := NewServer(int32(i), peers)
	// 	go server.Start()
	// }

	// cn := StartCentralNode() // khởi tạo node trung gian
	// log.Println(cn)

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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println(lis.Addr().Network())
	}

	srv := &Server{leaderId: int32(id), term: 1} // Tạo instance của server
	grpcServer := grpc.NewServer()
	raft.RegisterRaftServer(grpcServer, srv)

	// log.Println("number of peer in network %d", len(cn.nodes))
	// for index := 0; index < len(cn.nodes); index++ {
	// 	log.Println(cn.nodes[int32(index)])
	// }

	// go func() {
	// 	// Chờ 20 giây
	// 	time.Sleep(20 * time.Second)

	// 	// Bắt đầu bầu cử nếu không nhận được yêu cầu bầu cử nào
	// 	log.Printf("Node %d starting election after waiting 20 seconds", id)
	// 	srv.runElectionProcess() // Gọi runElectionProcess từ instance của server
	// }()

	// log.Printf("Node %d started on port %d", id, port)
	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	// Để giữ cho main goroutine không kết thúc
	select {}

}
