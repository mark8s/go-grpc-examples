package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/mark8s/go-grpc-examples/middleware/middleware/auth"
	"github.com/mark8s/go-grpc-examples/middleware/middleware/cred"
	"github.com/mark8s/go-grpc-examples/middleware/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type MiddlewareServer struct {
}

func (s *MiddlewareServer) Route(ctx context.Context, req *proto.SimpleRequest) (*proto.SimpleResponse, error) {
	res := proto.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	lis, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	grpcServer := grpc.NewServer(cred.TLSInterceptor(),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_auth.StreamServerInterceptor(auth.AuthInterceptor))),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor))))

	// 在gRPC服务器注册我们的服务
	proto.RegisterSimpleServer(grpcServer, &MiddlewareServer{})
	log.Println(Address + " net.Listing whth TLS and token...")
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
