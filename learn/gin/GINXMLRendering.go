package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/someXML", func(context *gin.Context) {
		// 方式一：自己拼接JSON
		context.XML(http.StatusOK, gin.H{
			"message": "someXML",
		})
	})
	engine.GET("/moreXML", func(context *gin.Context) {
		// 方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		context.XML(http.StatusOK, msg)
	})
	engine.Run()
}
