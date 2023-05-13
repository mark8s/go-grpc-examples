package main

import (
	"github.com/mark8s/go-grpc-examples/double-stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type DoubleStreamServer struct {
}

func (d *DoubleStreamServer) Conversations(req proto.Stream_ConversationsServer) error {

	for {
		recv, err := req.Recv()
		if err == io.EOF {
			return err
		}
		log.Println(recv)

		err = req.Send(&proto.StreamResponse{Answer: "I am mark!"})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// 创建tcp监听
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("listen err: " + err.Error())
	}

	grpcServer := grpc.NewServer()

	proto.RegisterStreamServer(grpcServer, &DoubleStreamServer{})

	log.Println("server running...")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("start grpc server err: " + err.Error())
	}

}
