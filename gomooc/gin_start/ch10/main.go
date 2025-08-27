package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	go func() {
		router.Run(":8085")
	}()

	// 如果想要接收到信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 监听Ctrl+C 和  kill
	<-quit                                               //只要quit接收到信号就会执行下面的代码

	fmt.Println("关闭server。。。")
}
