package main

import (
	"context"
	"github.com/mark8s/go-grpc-examples/server-side/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {

	// 创建连接
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 建立 grpc 连接，创建grpc client
	client := proto.NewStreamServerClient(conn)

	route, err := client.Route(ctx, &proto.SimpleRequest{Data: "grpc"})
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	log.Println(route)

	ctx2, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := client.ListValue(ctx2, &proto.SimpleRequest{
		Data: "stream server grpc",
	})
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}

		log.Println(recv.StreamValue)
	}

}
