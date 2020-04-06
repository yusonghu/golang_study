package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "someJSON",
			"key":     "helloWord",
		})
	})
	engine.GET("/moreJSON", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		context.JSON(http.StatusOK, msg)
	})
	engine.Run()
}
