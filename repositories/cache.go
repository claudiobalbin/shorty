package cache

import (
	"context"
	"log"
	"net/url"
	"regexp"
	"shorty/configs"

	"github.com/go-redis/redis/v8"
)

var settings = configs.GetSettings()

type CacheService struct {
	db *redis.Client
}

func NewCacheService() *CacheService {
	return &CacheService{
		db: redis.NewClient(&redis.Options{
			Addr:     settings["REDIS_URL"],
			Password: settings["REDIS_PASSWORD"],
			DB:       0,
		}),
	}
}

func (c *CacheService) GetUrl(key string) (string, bool) {
	if !validKey(key) {
		return "", false
	}

	val, err := c.db.Get(context.Background(), key).Result()
	if err != nil {
		log.Fatalf("Error getting key: %v", err)
	}

	return val, true
}

func (c *CacheService) SetUrl(key, longUrl string) bool {
	if !validKey(key) {
		return false
	}

	if !validURL(longUrl) {
		return false
	}

	err := c.db.Set(context.Background(), key, longUrl, 0).Err()
	if err != nil {
		log.Fatalf("Error setting key: %v", err)
	}

	return true
}

func validKey(key string) bool {
	if key == "" || len(key) < 3 {
		return false
	}

	return true
}

func validURL(urlStr string) bool {
	urlRegex := regexp.MustCompile(`^(http|https)://([\w-]+\.)+[\w-]+/`)

	if !urlRegex.MatchString(urlStr) {
		return false
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}
	if parsedURL.Host == "" {
		return false
	}
	if parsedURL.Path == "" {
		return false
	}

	return true
}
