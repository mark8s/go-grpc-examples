package main

import (
	"context"
	"github.com/mark8s/go-grpc-examples/deadline/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"runtime"
	"time"
)

type DeadlineServer struct {
}

func (ds *DeadlineServer) Route(ctx context.Context, req *proto.SimpleRequest) (*proto.SimpleResponse, error) {
	data := make(chan *proto.SimpleResponse, 1)
	go handle(ctx, req, data)
	select {
	case res := <-data:
		return res, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
}

func handle(ctx context.Context, req *proto.SimpleRequest, data chan<- *proto.SimpleResponse) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		runtime.Goexit() //超时后退出该Go协程
	case <-time.After(4 * time.Second): // 模拟耗时操作
		res := proto.SimpleResponse{
			Code:  200,
			Value: "hello " + req.Data,
		}
		// //修改数据库前进行超时判断
		// if ctx.Err() == context.Canceled{
		// 	...
		// 	//如果已经超时，则退出
		// }
		data <- &res
	}
}

var Address = ":8000"

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("listen err: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterSimpleServer(grpcServer, &DeadlineServer{})

	log.Printf(Address + " net.Listing...")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("start server err: %v", err)
	}

}
