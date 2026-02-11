package repository

import "errors"

var (
	ErrNotFound          = errors.New("record not found")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrDuplicateEntry    = errors.New("duplicate entry")
)
