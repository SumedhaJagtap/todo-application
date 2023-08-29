package utility

import "errors"

var (
	ErrRequestBodyEmpty = errors.New("empty request body provided")
	ErrTaskEmptyID      = errors.New("Task ID should be non empty")
	NoTaskFound         = errors.New("Task not found")
)
