package calculator

import "testing"

type fakeOperation struct {
	result float64
}
type fakeErrorOperation struct{}

func (fakeErrorOperation) Execute(operands []float64) (float64, error) {
	return 0, ErrDivisionByZero
}

func (f fakeOperation) Execute(operands []float64) (float64, error) {
	return f.result, nil
}

func TestCalculator_Calculate(t *testing.T) {
	calculator := NewCalculator(map[OperationType]Operation{
		Add: fakeOperation{result: 15},
	})

	result, err := calculator.Calculate(Add, []float64{10, 5})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != 15 {
		t.Errorf("expected result 15, got %v", result)
	}
}

func TestCalculator_InvalidOperation(t *testing.T) {
	calculator := NewCalculator(map[OperationType]Operation{})

	_, err := calculator.Calculate(Add, []float64{10, 5})

	if err != ErrInvalidOperation {
		t.Errorf("expected ErrInvalidOperation, got %v", err)
	}
}

func TestCalculator_PropagatesOperationError(t *testing.T) {
	calculator := NewCalculator(map[OperationType]Operation{
		Divide: fakeErrorOperation{},
	})

	_, err := calculator.Calculate(Divide, []float64{10, 0})

	if err != ErrDivisionByZero {
		t.Errorf("expected ErrDivisionByZero, got %v", err)
	}
}
