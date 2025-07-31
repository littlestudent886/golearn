package main

import (
	"context"
	"fmt"
	"gomooc/stream_grpc_test/proto"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//服务端流模式
	c := proto.NewGreeterClient(conn)

	res, _ := c.GetStream(context.Background(), &proto.Request{Data: "zzc"})
	for {
		r, err := res.Recv()
		if err != nil {
			break
		}
		log.Println(r.Data)
	}

	//客户端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		putS.Send(&proto.Request{Data: "zzc" + string(i)})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到服务器发送的数据" + data.Data)
			time.Sleep(time.Second * 1)
		}

	}()

	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.Request{Data: "AllStream客户端发送数据"})
			time.Sleep(time.Second * 1)
		}
	}()
	wg.Wait()

}
