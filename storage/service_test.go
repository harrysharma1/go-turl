package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testService = &Service{}

func init() {
	testService = InitialiseStorage()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testService.redisClient != nil)
}

func TestInsertionRetrieval(t *testing.T) {
	longUrl := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer"
	uuid := "4d67944c-72d6-4e3e-9dbc-850d92e05ac1"
	shortUrl := "1nahsuzinc"

	SaveUrlMapping(shortUrl, longUrl, uuid)

	retrievedUrl := GetInitialUrl(shortUrl)

	assert.Equal(t, longUrl, retrievedUrl)
}
