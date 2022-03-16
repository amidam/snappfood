package redis

import (
	"os"

	stdredis "github.com/go-redis/redis/v8"
)

type Redis struct {
	Redis *stdredis.Client
}

// New returns a naive implementation of DB.
func New() *Redis {
	return &Redis{
		Redis: Init(),
	}
}

func Init() *stdredis.Client {
	return stdredis.NewClient(
		&stdredis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		},
	)
}
