package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	//创建一个 HTTP 服务器
	return &http.Server{
		//监听地址和端口
		Addr: address,
		// 设置处理请求的处理程序
		Handler: router,
		//服务器读取请求的超时时间
		ReadTimeout: 20 * time.Second,
		//服务器写入响应的超时时间
		WriteTimeout: 20 * time.Second,
		//设置请求头的最大字节数,1mb
		MaxHeaderBytes: 1 << 20,
	}
}
