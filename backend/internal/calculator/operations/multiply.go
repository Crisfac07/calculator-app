package operations

import "calculator/backend/internal/calculator"

type MultiplyOperation struct{}

var _ calculator.Operation = MultiplyOperation{}

func (MultiplyOperation) Execute(operands []float64) (float64, error) {
	if len(operands) < 2 {
		return 0, calculator.ErrInvalidOperands
	}

	result := 1.0

	for _, operand := range operands {
		result *= operand
	}

	return result, nil
}
