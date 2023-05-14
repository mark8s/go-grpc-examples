package main

import (
	"context"
	"github.com/mark8s/go-grpc-examples/middleware/client/auth"
	"github.com/mark8s/go-grpc-examples/middleware/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {

	//从输入的证书文件中为客户端构造TLS凭证
	creds, err := credentials.NewClientTLSFromFile("../x509/ca.crt", "www.mark8s.com")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	//构建Token
	token := auth.Token{
		Value: "bearer grpc.auth.token",
	}
	// 连接服务器
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	client := proto.NewSimpleClient(conn)
	res, err := client.Route(context.Background(), &proto.SimpleRequest{
		Data: "grpc",
	})
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
