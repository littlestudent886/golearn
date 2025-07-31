package main

import (
	"context"
	"gomooc/helloworld/proto"
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
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	_ = s.Serve(listen)

}
