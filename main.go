package main

import (
	"fmt"
	"net/http"
	"turl/handler"
	"turl/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "goturl",
		})
	})

	router.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	router.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx)
	})

	storage.InitialiseStorage()

	err := router.Run(":6969")
	if err != nil {
		panic(fmt.Sprintf("Error occurring launching web server: %s", err))
	}
}
