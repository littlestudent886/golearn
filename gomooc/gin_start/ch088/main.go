package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", 123456)
		c.Next()
		latency := time.Since(t)
		fmt.Printf("latency:%v", latency)
		status := c.Writer.Status()
		fmt.Println("状态", status)
	}
}

func TokenRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		var token string
		for k, v := range context.Request.Header {
			if k == "X-Token" {
				token = v[0]
			}
		}
		if token != "zzc" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			context.Abort()
		}
		context.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(TokenRequired())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	router.Run(":8084")
}
