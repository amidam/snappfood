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

func (s *dispatcher) GetOrder(ctx context.Context, req requests.GetOrder) (responses.GetOrder, error) {
	r, err := json.Marshal(req)
	if err != nil {
		return responses.GetOrder{}, errors.Wrap(err, "service > GetOrder > Marshal")
	}

	if err = s.redisClient.RPush(ctx, "order", r).Err(); err != nil { // TODO: read from config
		return responses.GetOrder{}, errors.Wrap(err, "service > GetOrder > RPush")
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
