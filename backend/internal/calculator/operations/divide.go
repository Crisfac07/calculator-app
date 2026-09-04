package operations

import "calculator/backend/internal/calculator"

type DivideOperation struct{}

var _ calculator.Operation = DivideOperation{}

func (DivideOperation) Execute(operands []float64) (float64, error) {
	if len(operands) < 2 {
		return 0, calculator.ErrInvalidOperands
	}

	result := operands[0]

	for _, operand := range operands[1:] {
		if operand == 0 {
			return 0, calculator.ErrDivisionByZero
		}

		result /= operand
	}

	return result, nil
}
