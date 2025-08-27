package main

import (
	"context"
	"fmt"
	"gomooc/grpc_interpretor/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (res *proto.HelloReply, err error) {
	return &proto.HelloReply{Message: "hello" + req.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接收到了一个新的请求")
		res, err := handler(ctx, req)
		fmt.Println("请求已经完成")
		return res, err
	}

	opt := grpc.UnaryInterceptor(interceptor)
	s := grpc.NewServer(opt)
	proto.RegisterGreeterServer(s, &server{})
	_ = s.Serve(listen)

}
