package service

import (
	"context"

	"snappfood/A/pkg/db"
	"snappfood/A/pkg/redis"
)

type Process interface {
	ReadOrder(ctx context.Context) error
}

type process struct {
	redisClient *redis.Redis
	db          *db.DB
}

func NewProcess() Process {
	return &process{
		redisClient: redis.New(),
		db:          db.New(),
	}
}
