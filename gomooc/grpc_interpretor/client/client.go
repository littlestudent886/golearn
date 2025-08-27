package main

import (
	"context"
	"fmt"
	"gomooc/grpc_interpretor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("耗时:", time.Since(start))
		return err
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.NewClient(":50051", opts...)
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
