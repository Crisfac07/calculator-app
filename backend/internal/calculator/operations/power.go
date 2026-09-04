package operations

import (
	"calculator/backend/internal/calculator"
	"math"
)

type PowerOperation struct{}

var _ calculator.Operation = PowerOperation{}

func (PowerOperation) Execute(operands []float64) (float64, error) {
	if len(operands) != 2 {
		return 0, calculator.ErrInvalidOperands
	}

	return math.Pow(operands[0], operands[1]), nil
}
