package middleware

import (
	"snappfood/A/internal/dispatcher/service"

	"github.com/go-kit/log"
)

type Middleware func(service.Dispatcher) service.Dispatcher

func Logging(logger log.Logger) Middleware {
	return func(next service.Dispatcher) service.Dispatcher {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	service.Dispatcher
}