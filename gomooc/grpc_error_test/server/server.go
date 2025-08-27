package main

import (
	"context"
	"gomooc/grpc_error_test/proto"
	"google.golang.org/grpc"
	"net"
	"time"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (res *proto.HelloReply, err error) {
	time.Sleep(time.Second * 3)
	return &proto.HelloReply{Message: "hello" + req.Name}, nil
	//fmt.Println(codes.NotFound, "记录未找到:%s", req.Name)
	//return nil, status.Errorf(codes.NotFound, "记录未找到:%s", req.Name)
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
