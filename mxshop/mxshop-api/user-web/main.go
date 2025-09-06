package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"user-web/global"
	"user-web/initialize"
	"user-web/utils"

	myvalidator "user-web/validator"
)

func main() {
	//1.初始化日志
	initialize.InitLogger()

	//2.初始化配置
	initialize.InitConfig()

	//3.初始化routers
	Router := initialize.Routers()

	//4.初始化翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		return
	}

	//5.初始化Srv连接
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	// 如果是本地开发环境端口号固定，线上环境获取随机端口号
	debug := viper.GetBool("MXSHOP_DEBUG")
	zap.S().Info(debug)
	if !debug {
		port, err := utils.GetFreePort()
		if err != nil {
			return
		}
		global.ServerConfig.Port = port
	}
	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} ⾮法的⼿机号码!", true) // see universal-t
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

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
