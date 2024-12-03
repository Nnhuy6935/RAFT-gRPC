package main

// đây là node trung gian lưu thong tin các node trong mạng, nhận và trả về các phàn hổi giữa các node
import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "main/yourmodule/raft" // Thay thế "yourmodule" bằng module của bạn

	"google.golang.org/grpc"
)

var (
	centralNodeInstance *centralNode // Biến toàn cục để lưu instance của centralNode
	once                sync.Once    // Để đảm bảo rằng node chỉ được tạo một lần
)

func StartCentralNode() *centralNode {
	// conn, err := net.Dial("tcp","localhost:8002")
	// if err != nil
	// Sử dụng sync.Once để đảm bảo chỉ tạo một instance
	once.Do(func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8002")) // Địa chỉ cố định cho node trung gian
		if err != nil {
			// log.Fatalf("Failed to listen: %v", err)
			log.Println("Failed to listen: %v", err)
			return
		}

		grpcServer := grpc.NewServer()
		centralNodeInstance = &centralNode{nodes: make(map[int32]string)}

		pb.RegisterRaftServer(grpcServer, centralNodeInstance)

		log.Println("Central node listening on port 8002")
		go func() {
			if err := grpcServer.Serve(lis); err != nil {
				log.Fatalf("Failed to serve: %v", err)
			}
		}()
	})

	return centralNodeInstance // Trả về instance đã tạo
}

type centralNode struct {
	pb.UnimplementedRaftServer
	nodes      map[int32]string // Lưu trữ thông tin các node
	nodesMutex sync.Mutex       // Mutex để đồng bộ hóa truy cập
}

func (cn *centralNode) RegisterNode(ctx context.Context, req *pb.RegisterNodeRequest) (*pb.RegisterNodeResponse, error) {
	cn.nodesMutex.Lock()
	defer cn.nodesMutex.Unlock()

	cn.nodes[req.NodeId] = req.Address
	log.Printf("Node %d registered with address %s", req.NodeId, req.Address)

	return &pb.RegisterNodeResponse{Success: true}, nil
}

func (cn *centralNode) VoteRequest(ctx context.Context, req *pb.VoteRequest) (*pb.VoteResponse, error) {
	log.Printf("Received VoteRequest from candidate %d", req.CandidateId)

	// Logic xử lý bầu cử
	var voteGranted bool
	// Giả sử node trung gian đồng ý cấp phiếu
	voteGranted = true

	return &pb.VoteResponse{VoteGranted: voteGranted, Term: req.Term}, nil
}
