package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "GET",
		})
	})
	engine.POST("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "post",
		})
	})
	engine.PUT("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "put",
		})
	})
	engine.DELETE("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "delete",
		})
	})
	engine.Run()
}
