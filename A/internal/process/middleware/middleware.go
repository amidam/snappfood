package middleware

import (
	"snappfood/A/internal/process/service"

	"github.com/go-kit/log"
)

// middleware describes a service middleware.
type middleware func(service.Process) service.Process

func Logging(logger log.Logger) middleware {
	return func(next service.Process) service.Process {
		return &logging{
			next:   next,
			logger: logger,
		}
	}
}

type logging struct {
	next   service.Process
	logger log.Logger
}
