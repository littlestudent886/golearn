package main

import (
	"context"
	"fmt"
	"gomooc/metadata_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	// 设置metadata
	//metadata.Pairs("timestamp",time.Now().Format("2006-01-02 15:04:05"))
	md := metadata.New(map[string]string{
		"name": "zzc",
		"age":  "18",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	reply, err := c.SayHello(ctx,
		&proto.HelloRequest{
			Name: "zzc",
			Age:  "18",
			G:    proto.Gender_FEMALE,
			Mp: map[string]string{
				"name": "zzc",
				"age":  "18",
			},
			AddTime: timestamppb.Now(),
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply.Message)
}
