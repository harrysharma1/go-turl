package handler

import (
	"fmt"
	"net/http"
	"turl/storage"
	"turl/urlshortener"

	"github.com/gin-gonic/gin"
)

const host = "http://localhost:6969/"

type UrlModel struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var model UrlModel

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := urlshortener.ShortLinkGeneration(model.LongUrl, model.UserId)
	storage.SaveUrlMapping(shortUrl, model.LongUrl, model.UserId)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleAllRecentEntries(c *gin.Context) {
	res := storage.GetAllRecentUrlMappings()

	for key, val := range res {

		fmt.Printf("key: %s -> long url: %s\n -> short url: %s%s\n", key, val, host, key)
	}
	c.HTML(http.StatusOK, "index", gin.H{
		"title":  "Main website",
		"recent": res,
		"host":   host,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := storage.GetInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
