package cred

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// TLSInterceptor TLS证书认证
func TLSInterceptor() grpc.ServerOption {
	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	creds, err := credentials.NewServerTLSFromFile("../x509/server.crt", "../x509/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	return grpc.Creds(creds)
}
