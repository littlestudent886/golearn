package main

import (
	"fmt"
	"gomooc/stream_grpc_test/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct{}

func (s *server) GetStream(req *proto.Request, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.Response{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端发送的数据" + data.Data)
			time.Sleep(time.Second * 1)
		}

	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.Response{Data: "AllStream服务器端发送数据"})
			time.Sleep(time.Second * 1)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		return
	}
}
