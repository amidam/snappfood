package middleware

import (
	"context"
	"time"
)

func (mw *logging) ReadOrder(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "ReadOrder", "took", time.Since(begin))
		if err != nil {
			mw.logger.Log("err", err)
		}
	}(time.Now())
	return mw.next.ReadOrder(ctx)
}
