package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type Service struct {
	redisClient *redis.Client
}

var service = &Service{}

// var ctx = context.Background()

const CacheTimeout = 5 * time.Hour

func InitialiseStorage() *Service {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	res, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("error initialising redis: %v", err)
	}

	fmt.Printf("\nRedis started successfully\n--------------------------\nMessage: %s\n--------------------------", res)
	service.redisClient = client
	return service
}

func SaveUrlMapping(shortUrl string, longUrl, userId string) {
	cmd := service.redisClient.Set(shortUrl, longUrl, CacheTimeout)
	time := service.redisClient.Time()
	fmt.Println(time.String())
	if err := cmd.Err(); err != nil {
		log.Fatalf("error saving key: %v\n\nshortUrl: %s\nlongUrl: %s\n", err, shortUrl, longUrl)
	}
}

func GetInitialUrl(shortUrl string) string {
	res, err := service.redisClient.Get(shortUrl).Result()
	if err != nil {
		log.Fatalf("error retrieving initial url: %v\n shortUrl: %s\n", err, shortUrl)
	}
	return res
}

func GetTimeSinceCreation(key string) (time.Duration, error) {
	ttl, err := service.redisClient.TTL(key).Result()
	if err != nil {
		return 0, err
	}

	if ttl < 0 {
		return 0, fmt.Errorf("key has no expiration or does not exist")
	}

	return CacheTimeout - ttl, nil
}

func GetAllRecentUrlMappings() map[string]string {
	res := make(map[string]string)
	iter := service.redisClient.Scan(0, "*", 0).Iterator()

	for iter.Next() {
		key := iter.Val()
		val, err := service.redisClient.Get(key).Result()
		if err == nil {
			res[key] = val
		}
	}
	if err := iter.Err(); err != nil {
		log.Printf("error iterating redis keys: %s", err)
	}

	return res
}

func DeleteUrlMappingById(uuid string) {
	res := service.redisClient.Del(uuid)
	fmt.Println(res.Name())

}
