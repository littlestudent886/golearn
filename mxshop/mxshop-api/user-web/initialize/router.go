package initialize

import (
	"github.com/gin-gonic/gin"
	router "user-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/u/v1")
	// 用户路由
	router.InitUserRouter(ApiGroup)
	// 生成验证码路由
	router.InitBaseRouter(ApiGroup)
	return Router
}
