package calculator

import "errors"

var (
	ErrInvalidOperation   = errors.New("invalid operation")
	ErrInvalidOperands    = errors.New("invalid operands")
	ErrDivisionByZero     = errors.New("division by zero")
	ErrNegativeSquareRoot = errors.New("aquare root of negative number")
)
