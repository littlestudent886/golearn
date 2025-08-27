package main

import (
	"context"
	"fmt"
	"gomooc/grpc_error_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"time"
)

func main() {

	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	_, err = c.SayHello(ctx, &proto.HelloRequest{Name: " zzc"})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("解析error失败")
		}
		fmt.Println(st.Code())
		fmt.Println(st.Message())
	}
	//fmt.Println(reply.Message)
}
