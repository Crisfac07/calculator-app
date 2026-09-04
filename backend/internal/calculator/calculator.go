package calculator

type Calculator struct {
	operations map[OperationType]Operation
}

func NewCalculator(operations map[OperationType]Operation) *Calculator {
	return &Calculator{
		operations: operations,
	}
}

func (c *Calculator) Calculate(
	operationType OperationType,
	operands []float64,
) (float64, error) {
	operation, ok := c.operations[operationType]

	if !ok {
		return 0, ErrInvalidOperation
	}

	return operation.Execute(operands)
}
