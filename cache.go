package baseserver

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Client *redis.Client
}

func NewCache(cfg *Config) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisURL,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return &Cache{Client: client}
}
