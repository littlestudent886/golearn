package main

import (
	"context"
	"fmt"
	"gomooc/grpc_token_auth_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
		md, ok := metadata.FromIncomingContext(ctx)
		fmt.Println(md)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "未token认证")
		}
		var (
			appid  string
			appkey string
		)
		if val1, ok := md["appid"]; ok {
			appid = val1[0]
		}
		if val2, ok := md["appkey"]; ok {
			appkey = val2[0]
		}
		fmt.Println(appid, appkey)
		if appid != "10101" || appkey != "i am a key" {
			fmt.Println("token认证失败")
			return resp, status.Errorf(codes.Unauthenticated, "token认证失败")
		}
		res, err := handler(ctx, req)
		fmt.Println("请求已经完成")
		return res, err
	}

	opt := grpc.UnaryInterceptor(interceptor)
	s := grpc.NewServer(opt)
	proto.RegisterGreeterServer(s, &server{})
	_ = s.Serve(listen)

}
