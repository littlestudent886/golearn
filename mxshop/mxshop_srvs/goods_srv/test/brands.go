package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop_srvs/goods_srv/proto"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.NewClient("10.233.4.60:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	brandClient = proto.NewGoodsClient(conn)
}

func TestGetBrandList() {
	rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}
}

func TestCreateBrand() {
	rsp, err := brandClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: "测试",
		Logo: "http://www.baidu.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)
}

func TestDeleteBrand() {
	rsp, err := brandClient.DeleteBrand(context.Background(), &proto.BrandRequest{
		Id: 1113,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}
func TestUpdateBrand() {
	rsp, err := brandClient.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   1113,
		Name: "测试",
		Logo: "http://www.baidu.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
}

func main() {
	Init()
	defer conn.Close()
	//TestGetBrandList()
	//TestCreateBrand()
	TestUpdateBrand()
	//TestDeleteBrand()
}
