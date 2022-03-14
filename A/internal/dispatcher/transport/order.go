package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"snappfood/A/internal/dispatcher/endpoint"
)

func decodeGetOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.GetOrderRequest

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
