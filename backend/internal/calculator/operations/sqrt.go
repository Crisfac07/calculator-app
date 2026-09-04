package operations

import (
	"calculator/backend/internal/calculator"
	"math"
)

type SquareRootOperation struct{}

var _ calculator.Operation = SquareRootOperation{}

func (SquareRootOperation) Execute(operands []float64) (float64, error) {
	if len(operands) != 1 {
		return 0, calculator.ErrInvalidOperands
	}

	if operands[0] < 0 {
		return 0, calculator.ErrNegativeSquareRoot
	}

	return math.Sqrt(operands[0]), nil
}
