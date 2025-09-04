package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mxshop_srvs/user_srv/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName)
		checkRsp, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: user.Password,
		})
		if err != nil {
			return
		}
		fmt.Println(checkRsp.Success)
	}
	fmt.Println(rsp)
}

func TestCreateuser() {
	for i := 0; i < 10; i++ {
		userInfoRsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("zzc%d", i),
			Mobile:   fmt.Sprintf("1382568975%d", i),
			Password: "admin123",
		})
		if err != nil {
			return
		}
		fmt.Println(userInfoRsp.Id)
	}
}

func main() {
	Init()
	defer conn.Close()
	TestGetUserList()
	//TestCreateuser()
}
