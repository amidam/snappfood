package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"snappfood/A/internal/process/middleware"
	"snappfood/A/internal/process/service"

	"github.com/go-kit/log"
)

func Run() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var s service.Process
	{
		s = service.NewProcess()
		s = middleware.Logging(logger)(s)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	ctx := context.Background()
	go func() {
		logger.Log("Reading Order")
		errs <- s.ReadOrder(ctx)
	}()

	logger.Log("exit", <-errs)
}
