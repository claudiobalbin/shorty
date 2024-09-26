package repository

import (
	"context"
	"log"
	"shorty/configs"
	"shorty/utils"

	"github.com/go-redis/redis/v8"
)

var settings = configs.GetSettings()

type CacheRepository struct {
	db *redis.Client
}

func NewCacheRepository() *CacheRepository {
	return &CacheRepository{
		db: redis.NewClient(&redis.Options{
			Addr:     settings["REDIS_URL"],
			Password: settings["REDIS_PASSWORD"],
			DB:       0,
		}),
	}
}

func (c *CacheRepository) GetUrl(key string) (string, bool) {
	if !utils.ValidKey(key) {
		return "", false
	}

	val, err := c.db.Get(context.Background(), key).Result()
	if err != nil {
		log.Fatalf("Error getting key: %v", err)
	}

	return val, true
}

func (c *CacheRepository) SetUrl(key, longUrl string) bool {
	if !utils.ValidKey(key) {
		return false
	}

	if !utils.ValidURL(longUrl) {
		return false
	}

	err := c.db.Set(context.Background(), key, longUrl, 0).Err()
	if err != nil {
		log.Fatalf("Error setting key: %v", err)
	}

	return true
}

func (c *CacheRepository) CleanCache(ctx context.Context) (string, bool) {
	val, err := c.db.FlushAll(ctx).Result()
	if err != nil {
		log.Fatalf("Error cleaning cache: %v", err)
	}

	return val, true
}
