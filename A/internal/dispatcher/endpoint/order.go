package endpoint

import (
	"context"

	"snappfood/A/internal/dispatcher/model"
	"snappfood/A/internal/dispatcher/model/requests"
	"snappfood/A/internal/dispatcher/model/responses"
	"snappfood/A/internal/dispatcher/service"

	kitendpoint "github.com/go-kit/kit/endpoint"
)

func MakeGetOrder(svc service.Dispatcher) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(GetOrderRequest); ok {
			resp, err := svc.GetOrder(ctx, req.GetOrder)
			return GetOrderResponse{resp, err}, nil
		}
		return nil, model.ErrTypeAssertion
	}
}

type GetOrderRequest struct {
	requests.GetOrder
}

type GetOrderResponse struct {
	responses.GetOrder
	Err error `json:"err,omitempty"`
}

func (r GetOrderResponse) Failed() error { return r.Err }
