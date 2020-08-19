package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()



	router.GET("/ping", func(context *gin.Context) {
		response := make(map[string]string)
		response["message"] = "hi"
		context.JSON(200, response)
	})
	router.Run(":8081")
}