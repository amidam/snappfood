package cmd

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"snappfood/A/internal/dispatcher/middleware"
	"snappfood/A/internal/dispatcher/service"
	"snappfood/A/internal/dispatcher/transport"

	"github.com/go-kit/log"
)

func Run() {
	httpAddr := flag.String("addr", ":8080", "HTTP listen address")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var s service.Dispatcher
	{
		s = service.NewDispatcher()
		s = middleware.Logging(logger)(s)
	}

	var h http.Handler
	{
		h = transport.MakeHandler(s, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
