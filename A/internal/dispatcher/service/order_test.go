package service

import (
	"context"
	"testing"

	"snappfood/A/internal/dispatcher/model/requests"
)

func TestGetOrderErrorResponse(t *testing.T) {
	d := NewDispatcher()

	req := requests.GetOrder{
		ID:    10,
		Price: 1000,
		Title: "burger",
	}

	_, err := d.GetOrder(context.Background(), req)
	if err != nil {
		t.Fatalf("got %v, want %v", err, nil)
	}
}

func TestGetOrderResponse(t *testing.T) {
	d := NewDispatcher()

	req := requests.GetOrder{
		ID:    100,
		Price: 10000,
		Title: "burger2",
	}

	resp, _ := d.GetOrder(context.Background(), req)
	if resp.ID != req.ID {
		t.Fatalf("got %v, want %v", req.ID, resp.ID)
	}
}
