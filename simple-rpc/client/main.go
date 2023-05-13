package main

import (
	"context"
	"fmt"
	"github.com/mark8s/go-grpc-examples/simple-rpc/helloworld"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	client := helloworld.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	search, err := client.Search(ctx, &helloworld.HelloRequest{
		Name: "mark",
	})
	if err != nil {
		return
	}
	fmt.Println(search)
}
