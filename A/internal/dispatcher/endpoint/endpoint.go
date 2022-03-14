package endpoint

import (
	"snappfood/A/internal/dispatcher/service"

	kitendpoint "github.com/go-kit/kit/endpoint"
)

type set struct {
	GetOrderEndpoint kitendpoint.Endpoint
}

func MakeServer(s service.Dispatcher) set {
	return set{
		GetOrderEndpoint: MakeGetOrder(s),
	}
}
