package main

import (
	"context"
	"github.com/mark8s/go-grpc-examples/security/pkg/auth"
	"github.com/mark8s/go-grpc-examples/security/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile("../pkg/x509/ca.crt", "www.mark8s.com")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	//构建Token
	token := auth.Token{
		AppID:     "grpc_token",
		AppSecret: "123456",
	}

	// 连接服务器
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	client := proto.NewSimpleClient(conn)
	route, err := client.Route(context.Background(), &proto.SimpleRequest{})
	if err != nil {
		log.Fatalf("call route err: %v", err)
	}

	log.Printf(route.String())
}
