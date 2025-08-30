package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	err := router.Run(":6969")
	if err != nil {
		panic(fmt.Sprintf("Error occurring launching web server: %s", err))
	}
}
