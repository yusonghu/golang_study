package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/someYaml", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{
			"message": "someYaml",
		})
	})
	engine.GET("/moreYaml", func(context *gin.Context) {
		type user struct {
			Name string
			Age  int8
		}
		u := user{
			"张三",
			23,
		}
		context.YAML(http.StatusOK, u)
	})
	engine.Run()
}
