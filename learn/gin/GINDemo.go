package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	router := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello", "world")
	})
	router.PUT("/put", func(context *gin.Context) {
		context.JSON(200, "已经更新了数据")
	})
	router.DELETE("/delete", func(context *gin.Context) {
		context.JSON(200, "数据已经被删除")
	})
	router.POST("/post", func(context *gin.Context) {
		key := context.Param("key")
		fmt.Println(key)
		context.JSON(200, "已收到数据"+key)
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	router.Run()
}
