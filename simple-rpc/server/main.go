package main

import (
	"context"
	pb "github.com/mark8s/go-grpc-examples/simple-rpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloServer struct {
}

func (s *HelloServer) Search(ctx context.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", r.GetName())
	return &pb.HelloResponse{Message: "Hello " + r.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	// 注册 helloworld 服务到 grpc服务器中
	pb.RegisterHelloServiceServer(grpcServer, &HelloServer{})

	log.Printf("server listening at %v", lis.Addr())
	// 启动grpc 服务器
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
