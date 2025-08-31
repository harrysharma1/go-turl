package main

import (
	"fmt"
	"turl/handler"
	"turl/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		handler.HandleAllRecentEntries(ctx)
	})

	router.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	router.POST("/delete-short-url/:uuid", func(ctx *gin.Context) {
		handler.HandleDeleteUrlById(ctx)
	})

	router.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx)
	})

	router.GET("/recent-urls", handler.HandleAllRecentEntriesJson)

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // No Content, or serve a real favicon if you want
	})
	storage.InitialiseStorage()

	err := router.Run(":6969")
	if err != nil {
		panic(fmt.Sprintf("Error occurring launching web server: %s", err))
	}
}
