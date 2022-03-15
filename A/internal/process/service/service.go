package service

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Process interface {
	ReadOrder(ctx context.Context) error
}

type process struct {
	redisClient *redis.Client
}

func NewDispatcher() Process {
	return &process{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // TODO: read from config
			Password: "",
			DB:       0,
		}),
	}
}
