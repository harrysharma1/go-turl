package main

import (
	"fmt"
	"turl/handler"
	"turl/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to URL shortener.",
		})
	})

	router.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	router.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx)
	})

	storage.InitialiseStorage()

	err := router.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Error occurring launching web server: %s", err))
	}
}
