package main

import (
	"context"
	"fmt"
	proto_bak "gomooc/grpc_proto_test/proto-bak"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	c := proto_bak.NewGreeterClient(conn)
	reply, err := c.SayHello(context.Background(),
		&proto_bak.HelloRequest{
			Name: "zzc",
			Age:  "18",
			G:    proto_bak.Gender_FEMALE,
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
