package redisclient

import (
	"context"
	"github.com/redis/go-redis/v9"
	"ms-shopify/pkg/logster"
	"sync"
)

var (
	rdb  *redis.Client
	once sync.Once
	ctx  = context.Background()
)

// Config holds the Redis configuration.
type Config struct {
	Addr     string
	Password string
	DB       int
}

func SetRedisClient(cfg Config) {
	logster.Info("Setting Redis client")
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB,
		})
	})
	logster.Info("Redis client set")
}

// GetContext returns the Redis context.
func GetContext() context.Context {
	return ctx
}

func GetRedisClient() *redis.Client {
	return rdb
}
