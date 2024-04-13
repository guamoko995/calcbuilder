package calcbuilder

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidTerm                        = errors.New("invalid term")
	ErrUnexpectedEndOfExpression          = errors.New("anexpected end of expression")
	ErrUnexpectedContinuationOfExpression = errors.New("unexpected continuation of expression")
)

// PositionErr wraps the error and stores the problematic position in the expression.
type PositionErr struct {
	Position   int
	WrappedErr error
}

func (err PositionErr) Unwrap() error {
	return err.WrappedErr
}

func (err PositionErr) Error() string {
	return fmt.Sprintf("%s at position %d", err.WrappedErr.Error(), err.Position)
}
