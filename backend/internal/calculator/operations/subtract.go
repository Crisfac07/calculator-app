package operations

import "calculator/backend/internal/calculator"

type SubtractOperation struct{}

var _ calculator.Operation = SubtractOperation{}

func (SubtractOperation) Execute(operands []float64) (float64, error) {
	if len(operands) < 2 {
		return 0, calculator.ErrInvalidOperands
	}

	result := operands[0]

	for _, operand := range operands[1:] {
		result -= operand
	}

	return result, nil
}
