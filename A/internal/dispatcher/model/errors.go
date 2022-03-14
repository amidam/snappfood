package model

import "github.com/pkg/errors"

var (
	ErrUnexpected    = errors.New("unexpected error")
	ErrTypeAssertion = errors.New("invalid type assertion")
)
