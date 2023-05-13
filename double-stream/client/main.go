package main

import (
	"context"
	"github.com/mark8s/go-grpc-examples/double-stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync"
	"time"
)

func main() {

	// 创建连接
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("conn err " + err.Error())
	}

	// 创建客户端
	client := proto.NewStreamClient(conn)
	stream, err := client.Conversations(context.Background())
	if err != nil {
		log.Fatalf("call err " + err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			err := stream.Send(&proto.StreamRequest{Question: "what is your name?"})
			if err != nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	go func() {
		for {
			recv, err := stream.Recv()
			if err == io.EOF {
				stream.CloseSend()
				break
			}
			log.Println(recv.Answer)
		}
		wg.Done()
	}()
	wg.Wait()

}
