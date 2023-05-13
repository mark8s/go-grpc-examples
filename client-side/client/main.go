package main

import (
	"context"
	"github.com/mark8s/go-grpc-examples/client-side/proto"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

func main() {
	// 创建连接
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("conn grpc server err" + err.Error())
	}

	// 初始化grpc 客户端
	client := proto.NewStreamClientClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := client.RouteList(ctx)
	if err != nil {
		log.Fatalf("call grpc server err" + err.Error())
	}

	for i := 0; i < 5; i++ {
		err := stream.Send(&proto.StreamRequest{
			StreamData: "stream client rpc " + strconv.Itoa(i),
		})
		if err != nil {
			log.Fatalf("send data err" + err.Error())
		}
	}

	// 关闭流并接受响应
	recv, err := stream.CloseAndRecv()
	if err != nil {
		return
	}

	log.Println(recv)
}
