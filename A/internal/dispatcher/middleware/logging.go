package middleware

import (
	"snappfood/A/internal/dispatcher/service"

	"github.com/go-kit/log"
)

// middleware describes a service middleware.
type middleware func(service.Dispatcher) service.Dispatcher

func Logging(logger log.Logger) middleware {
	return func(next service.Dispatcher) service.Dispatcher {
		return &logging{
			next:   next,
			logger: logger,
		}
	}
}

type logging struct {
	next   service.Dispatcher
	logger log.Logger
}
