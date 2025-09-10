package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/handler"
	"mxshop_srvs/goods_srv/initialize"
	"mxshop_srvs/goods_srv/proto"
	"mxshop_srvs/goods_srv/utils"
)

func main() {
	IP := flag.String("ip", "10.233.4.60", "ip地址")
	Port := flag.Int("port", 0, "端口号")

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()
	zap.S().Infof("ip:", *IP)
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Infof("port:", *Port)

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 注册健康服务
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = global.ServerConfig.Name + serviceId
	registration.Address = "10.233.4.60"
	registration.Port = *Port
	registration.Tags = []string{"imooc", "zzc"}

	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("10.233.4.60:%d", *Port),
		Interval:                       "5s",
		Timeout:                        "3s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration.Check = check
	// 启动两个服务
	// 注册到consul不被覆盖
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceId); err != nil {
		zap.S().Info("[Consul]注销服务失败")
	}
	zap.S().Info("[Consul]注销服务成功")
}
