package errors

import "fmt"

var (
	ErrInvalidMove = fmt.Errorf("invalid move: cell is already occupied")
	ErrOutOfBounds = fmt.Errorf("invalid move: coordinates out of bounds")
	ErrInvalidInput = fmt.Errorf("invalid input")
	ErrUnknown = fmt.Errorf("unknown error occurred")
)