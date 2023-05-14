package main

import (
	"context"
	"github.com/mark8s/go-grpc-examples/security/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type SecurityServer struct {
}

func (s *SecurityServer) Route(ctx context.Context, req *proto.SimpleRequest) (*proto.SimpleResponse, error) {
	return &proto.SimpleResponse{
		Code:  200,
		Value: "hello",
	}, nil
}

var Address = ":8000"

func main() {

	// 监听本地端口
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	creds, err := credentials.NewServerTLSFromFile("../pkg/x509/server.crt", "../pkg/x509/server.key")
	if err != nil {
		log.Fatalf("Failed to generate TLS credentials %v", err)
	}

	//普通方法：一元拦截器（grpc.UnaryInterceptor）
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//拦截普通方法请求，验证Token
		err = Check(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	// 新建gRPC服务器实例,并开启TLS认证
	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))

	proto.RegisterSimpleServer(grpcServer, &SecurityServer{})

	log.Println(Address + " net.Listing whth TLS and token...")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func Check(ctx context.Context) error {
	// 从上下文中取数据
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return status.Errorf(codes.Unauthenticated, "获取Token失败")
	}
	var (
		appID     string
		appSecret string
	)

	if value, ok := md["app_id"]; ok {
		appID = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}

	if appID != "grpc_token" || appSecret != "123456" {
		return status.Errorf(codes.Unauthenticated, "Token无效: app_id=%s, app_secret=%s", appID, appSecret)
	}
	return nil
}
