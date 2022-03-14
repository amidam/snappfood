package service

import (
	"context"

	"snappfood/A/internal/dispatcher/model/requests"
	"snappfood/A/internal/dispatcher/model/responses"

	"github.com/go-redis/redis/v8"
)

type Dispatcher interface {
	GetOrder(ctx context.Context, req requests.GetOrder) (responses.GetOrder, error)
}

type dispatcher struct {
	redisClient *redis.Client
}

func NewDispatcher() Dispatcher {
	return &dispatcher{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // TODO: read from config
			Password: "",
			DB:       0,
		}),
	}
}
