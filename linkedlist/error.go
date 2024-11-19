package linkedlist

import "errors"

var (
	ErrInvalidPos = errors.New("Invalid position")
	ErrNotFound   = errors.New("Not found")
	ErrEmptyList  = errors.New("Empty list")
)
