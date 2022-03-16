package service

import (
	"context"
	"encoding/json"
	"net/http"

	"snappfood/A/internal/dispatcher/model"
	"snappfood/A/internal/dispatcher/model/requests"
	"snappfood/A/internal/dispatcher/model/responses"

	"github.com/pkg/errors"
)

func (d *dispatcher) GetOrder(ctx context.Context, req requests.GetOrder) (responses.GetOrder, error) {
	r, err := json.Marshal(req)
	if err != nil {
		return responses.GetOrder{}, errors.Wrap(err, "Marshal")
	}

	if err = d.redisClient.RPush(ctx, "order", r).Err(); err != nil {
		return responses.GetOrder{}, errors.Wrap(err, "RPush")
	}

	return responses.GetOrder{
		ID: req.ID,
		General: responses.General{
			Description: model.MsgSuccess,
			Status:      http.StatusOK,
			Code:        model.CodeSuccess,
		},
	}, nil
}
