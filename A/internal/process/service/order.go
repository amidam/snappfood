package service

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func (p *process) ReadOrder(ctx context.Context) error {
	for {
		result, err := p.redisClient.Redis.BLPop(ctx, 0*time.Second, "order").Result()
		if err != nil {
			return errors.Wrap(err, "BLPop")
		}

		params := map[string]interface{}{}
		err = json.NewDecoder(strings.NewReader(result[1])).Decode(&params)
		if err != nil {
			return errors.Wrap(err, "json decoder")
		}

		orderID, err := p.db.SaveOrder(params)
		if err != nil {
			return err
		}
		log.Printf("Order ID %s processed successfully.", orderID)
	}
}
