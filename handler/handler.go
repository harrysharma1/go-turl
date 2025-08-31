package handler

import (
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

func CreateShortUrl(ctx *gin.Context) {
	// var model UrlModel

	longUrl := ctx.Request.FormValue("longUrl")
	uid := ctx.Request.FormValue("uid")

	shortUrl := urlshortener.ShortLinkGeneration(longUrl, uid)
	storage.SaveUrlMapping(shortUrl, longUrl, uid)

	ctx.Redirect(http.StatusFound, "/")
}

func HandleAllRecentEntries(ctx *gin.Context) {
	res := storage.GetAllRecentUrlMappings()

	ctx.HTML(http.StatusOK, "index", gin.H{
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
