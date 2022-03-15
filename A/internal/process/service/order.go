package service

import (
	"context"
	"time"
)

func (p *process) ReadOrder(ctx context.Context) error {
	for {
		result, err := p.redisClient.BLPop(ctx, 0*time.Second, "order").Result() // TODO: read from config
	}
}
