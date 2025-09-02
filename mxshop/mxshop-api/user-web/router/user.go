package router

import (
	"github.com/gin-gonic/gin"
	"user-web/api"
	middlewares "user-web/midddlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
	}
}
