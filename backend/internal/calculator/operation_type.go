package calculator

type OperationType string

const (
	Add        OperationType = "add"
	Subtract   OperationType = "subtract"
	Multiply   OperationType = "multiply"
	Divide     OperationType = "divide"
	Power      OperationType = "power"
	SquareRoot OperationType = "sqrt"
)
