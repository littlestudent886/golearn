package main

import (
	"fmt"
	"go.uber.org/zap"
	"user-web/global"
	"user-web/initialize"
)

func main() {
	//1.初始化日志
	initialize.InitLogger()

	//2.初始化配置
	initialize.InitConfig()

	//3.初始化routers
	Router := initialize.Routers()

	/*
		1.S()可以获取一个全局的sugar，可以自己设置一个全局的logger
		2.日志是分级别的，debug，info，warn，error，panic，fatal
		3.S函数和L函数很有用，可以提供一个全局的安全访问logger的途径
	*/
	zap.S().Info("启动服务器，端口：", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
