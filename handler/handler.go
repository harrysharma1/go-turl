package handler

import (
	"net/http"
	"turl/storage"
	"turl/urlshortener"

	"github.com/gin-gonic/gin"
)

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

	host := "http://localhost:6969/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := storage.GetInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
