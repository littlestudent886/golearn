package main

import (
	"context"
	"fmt"
	"gomooc/helloworld/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	reply, err := c.SayHello(context.Background(), &proto.HelloRequest{})
	if err != nil {
		return
	}
	fmt.Println(reply.Message)
}
