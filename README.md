# Lab1 - Blockchain : RAFT Algorithm Implementation 
- download và install Consul: https://developer.hashicorp.com/consul/install
- Khởi động Consul trước khi build:  `consul agent -dev`
- web html consul : http://localhost:8500/ui/dc1/nodes
- biên dịch proto: `protoc --go_out=. --go-grpc_out=. raft.proto`
