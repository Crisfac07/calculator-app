package operations

import "calculator/backend/internal/calculator"

type AddOperation struct{}

var _ calculator.Operation = AddOperation{}

func (AddOperation) Execute(operands []float64) (float64, error) {
	if len(operands) < 2 {
		return 0, calculator.ErrInvalidOperands
	}

	var result float64

	for _, operand := range operands {
		result += operand
	}

	return result, nil
}
