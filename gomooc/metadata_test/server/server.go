package main

import (
	"context"
	"fmt"
	"gomooc/metadata_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (res *proto.HelloReply, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata error")
	}
	//for key, value := range md {
	//	fmt.Println(key, value)
	//}

	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}

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
