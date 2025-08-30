package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	redisClient *redis.Client
}

var service = &Service{}
var ctx = context.Background()

const CacheTimeout = 2 * time.Hour

func InitialiseStorage() *Service {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6969",
		Password: "",
		DB:       0,
	})

	res, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("error initialising redis: %v", err)
	}

	fmt.Printf("\nRedis started successfully\n--------------------------\nMessage: %s\n--------------------------", res)
	service.redisClient = client
	return service
}

func SaveUrlMapping(shortUrl string, longUrl, userId string) {
	err := service.redisClient.Set(ctx, shortUrl, longUrl, CacheTimeout)
	if err != nil {
		log.Fatalf("error saving key: %v\n shortUrl:%s\nlongUrl:%s\n", err, shortUrl, longUrl)
	}
}

func GetInitialUrl(shortUrl string) string {
	res, err := service.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		log.Fatalf("error retrieving initial url: %v\n shortUrl: %s\n", err, shortUrl)
	}
	return res
}
