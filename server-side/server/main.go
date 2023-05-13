package main

import (
	"context"
	pb "github.com/mark8s/go-grpc-examples/server-side/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type StreamServer struct {
}

func (c *StreamServer) Route(ctx context.Context, r *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + r.Data,
	}
	return &res, nil
}

func (c *StreamServer) ListValue(r *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&pb.StreamResponse{
			StreamValue: r.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStreamServerServer(grpcServer, &StreamServer{})

	log.Printf("server listening at %v", lis.Addr())
	// 启动grpc 服务器
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
