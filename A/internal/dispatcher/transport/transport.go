package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"snappfood/A/internal/dispatcher/endpoint"
	"snappfood/A/internal/dispatcher/model"
	"snappfood/A/internal/dispatcher/model/responses"
	"snappfood/A/internal/dispatcher/service"

	kitendpoint "github.com/go-kit/kit/endpoint"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func MakeHandler(svc service.Dispatcher, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := endpoint.MakeServer(svc)
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods(http.MethodPost).Path("/api/order").Handler(kithttp.NewServer( // TODO
		e.GetOrderEndpoint,
		decodeGetOrderRequest,
		encodeResponse,
		options...,
	))
	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(kitendpoint.Failer); ok && e.Failed() != nil {
		encodeError(ctx, e.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errToHTTPStatus(err))
	json.NewEncoder(w).Encode(
		responses.General{
			Description: errToEnglish(err).Error(),
			Status:      errToHTTPStatus(err),
			Code:        errToCode(err),
		},
	)
}

func errToEnglish(err error) error {
	switch errors.Cause(err) {
	case model.ErrTypeAssertion:
		return err
	default:
		return model.ErrUnexpected
	}
}

func errToHTTPStatus(err error) int {
	switch errors.Cause(err) {
	default:
		return http.StatusInternalServerError
	}
}

func errToCode(err error) int {
	switch errors.Cause(err) {
	case model.ErrTypeAssertion:
		return model.CodeErrTypeAssertion
	default:
		return model.CodeErrUnexpected
	}
}
