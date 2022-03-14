package middleware

import (
	"context"
	"time"

	"snappfood/A/internal/dispatcher/model/requests"
	"snappfood/A/internal/dispatcher/model/responses"
)

func (mw logmw) GetOrder(ctx context.Context, req requests.GetOrder) (resp responses.GetOrder, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetOrder",
			"id", req.ID,
			"took", time.Since(begin),
		)
		if err != nil {
			mw.logger.Log("err", err)
		}
	}(time.Now())

	resp, err = mw.Dispatcher.GetOrder(ctx, req)
	return
}
