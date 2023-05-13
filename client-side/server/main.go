package main

import (
	"github.com/mark8s/go-grpc-examples/client-side/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type StreamClientServer struct {
}

func (c *StreamClientServer) RouteList(stream proto.StreamClient_RouteListServer) error {
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			log.Println(err)
			// 发送结果，并关闭
			return stream.SendAndClose(&proto.SimpleResponse{Value: "ok"})
		}

		if err != nil {
			return err
		}
		log.Println(recv.StreamData)
	}
	return nil
}

func main() {
	// 创建tcp监听
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf(err.Error())
	}

	grpcServer := grpc.NewServer()
	// 注册server到grpc server
	proto.RegisterStreamClientServer(grpcServer, &StreamClientServer{})

	// 启动grpc server
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
