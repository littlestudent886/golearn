package main

import (
	"context"
	"fmt"
	"gomooc/grpc_token_auth_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type customCredentials struct {
}

func (c customCredentials) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "10101",
		"appkey": "i am a key",
	}, nil
}
func (c customCredentials) RequireTransportSecurity() bool {
	return false
}
func main() {
	//interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	start := time.Now()
	//	md := metadata.New(map[string]string{
	//		"appid":  "1010",
	//		"appkey": "i am a key",
	//	})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Println("耗时:", time.Since(start))
	//	return err
	//}

	interceptor := grpc.WithPerRPCCredentials(customCredentials{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	opts = append(opts, interceptor)
	conn, err := grpc.NewClient(":50051", opts...)
	if err != nil {
		return
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	reply, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "zzc"})
	if err != nil {
		return
	}
	fmt.Println(reply.Message)
}
