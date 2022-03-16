package service

import (
	"context"
	"testing"
)

func TestReadOrderErrorResponse(t *testing.T) {
	d := NewProcess()

	err := d.ReadOrder(context.Background())
	if err != nil {
		t.Fatalf("got %v, want %v", err, nil)
	}
}
