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

	ctx.JSON(http.StatusOK, gin.H{
		"status": "created",
	})
}

func HandleAllRecentEntries(ctx *gin.Context) {
	res := storage.GetAllRecentUrlMappings()

	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":  "Main website",
		"recent": res,
		"host":   host,
	})

}

func HandleAllRecentEntriesJson(ctx *gin.Context) {
	res := storage.GetAllRecentUrlMappings()
	type UrlEntry struct {
		Key     string `json:"key"`
		Value   string `json:"value"`
		Elapsed string `json:"elapsed"`
	}
	var urlEntries []UrlEntry
	for key, value := range res {
		elapsed, err := storage.GetTimeSinceCreation(key)
		elapsedStr := ""
		if err == nil {
			elapsedStr = elapsed.String()
		}
		urlEntries = append(urlEntries, UrlEntry{
			Key:     key,
			Value:   value,
			Elapsed: elapsedStr,
		})
	}
	ctx.JSON(http.StatusOK, urlEntries)
}

func HandleDeleteUrlById(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	storage.DeleteUrlMappingById(uuid)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}

func HandleShortUrlRedirect(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	initialUrl := storage.GetInitialUrl(shortUrl)
	ctx.Redirect(302, initialUrl)
}
