package calculator

import (
	"calculator/backend/internal/calculator"
)

func errorCode(err error) string {
	switch err {
	case calculator.ErrInvalidOperation:
		return "INVALID_OPERATION"
	case calculator.ErrInvalidOperands:
		return "INVALID_OPERANDS"
	case calculator.ErrDivisionByZero:
		return "DIVISION_BY_ZERO"
	case calculator.ErrNegativeSquareRoot:
		return "NEGATIVE_SQUARE_ROOT"
	default:
		return "INTERNAL_ERROR"
	}
}
