package operations

import (
	"calculator/backend/internal/calculator"
	"testing"
)

func TestAddOperation_Execute(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		expected float64
	}{
		{
			name:     "two operands",
			operands: []float64{10, 5},
			expected: 15,
		},
		{
			name:     "multiple operands",
			operands: []float64{10, 5, 3},
			expected: 18,
		},
		{
			name:     "negative numbers",
			operands: []float64{-10, 5},
			expected: -5,
		},
		{
			name:     "decimal numbers",
			operands: []float64{10.5, 2.5},
			expected: 13,
		},
	}

	operation := AddOperation{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := operation.Execute(test.operands)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result != test.expected {
				t.Errorf("expected result %v, got %v", test.expected, result)
			}
		})
	}
}

func TestAddOperation_Execute_InvalidOperands(t *testing.T) {
	operation := AddOperation{}

	_, err := operation.Execute([]float64{10})

	if err != calculator.ErrInvalidOperands {
		t.Errorf("expected ErrInvalidOperands, got %v", err)
	}
}

func TestSubtractOperation_Execute(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		expected float64
	}{
		{
			name:     "two operands",
			operands: []float64{10, 5},
			expected: 5,
		},
		{
			name:     "multiple operands",
			operands: []float64{100, 20, 10},
			expected: 70,
		},
		{
			name:     "negative result",
			operands: []float64{5, 10},
			expected: -5,
		},
		{
			name:     "negative numbers",
			operands: []float64{-10, -5},
			expected: -5,
		},
		{
			name:     "decimal numbers",
			operands: []float64{10.5, 2.5},
			expected: 8,
		},
	}

	operation := SubtractOperation{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := operation.Execute(test.operands)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result != test.expected {
				t.Errorf("expected result %v, got %v", test.expected, result)
			}
		})
	}
}

func TestSubtractOperation_Execute_InvalidOperands(t *testing.T) {
	operation := SubtractOperation{}

	_, err := operation.Execute([]float64{10})

	if err != calculator.ErrInvalidOperands {
		t.Errorf("expected ErrInvalidOperands, got %v", err)
	}
}

func TestMultiplyOperation_Execute(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		expected float64
	}{
		{
			name:     "two operands",
			operands: []float64{10, 5},
			expected: 50,
		},
		{
			name:     "multiple operands",
			operands: []float64{2, 3, 4},
			expected: 24,
		},
		{
			name:     "negative numbers",
			operands: []float64{-10, 5},
			expected: -50,
		},
		{
			name:     "zero",
			operands: []float64{10, 5, 0},
			expected: 0,
		},
		{
			name:     "decimal numbers",
			operands: []float64{2.5, 4},
			expected: 10,
		},
	}

	operation := MultiplyOperation{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := operation.Execute(test.operands)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result != test.expected {
				t.Errorf("expected result %v, got %v", test.expected, result)
			}
		})
	}
}

func TestMultiplyOperation_Execute_InvalidOperands(t *testing.T) {
	operation := MultiplyOperation{}

	_, err := operation.Execute([]float64{10})

	if err != calculator.ErrInvalidOperands {
		t.Errorf("expected ErrInvalidOperands, got %v", err)
	}
}

func TestDivideOperation_Execute(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		expected float64
	}{
		{
			name:     "two operands",
			operands: []float64{10, 2},
			expected: 5,
		},
		{
			name:     "multiple operands",
			operands: []float64{100, 2, 5},
			expected: 10,
		},
		{
			name:     "negative result",
			operands: []float64{-10, 2},
			expected: -5,
		},
		{
			name:     "decimal numbers",
			operands: []float64{10.5, 2},
			expected: 5.25,
		},
	}

	operation := DivideOperation{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := operation.Execute(test.operands)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result != test.expected {
				t.Errorf("expected result %v, got %v", test.expected, result)
			}
		})
	}
}

func TestDivideOperation_Execute_DivisionByZero(t *testing.T) {
	tests := [][]float64{
		{10, 0},
		{100, 2, 0},
	}

	operation := DivideOperation{}

	for _, operands := range tests {
		_, err := operation.Execute(operands)

		if err != calculator.ErrDivisionByZero {
			t.Errorf(
				"expected ErrDivisionByZero, got %v",
				err,
			)
		}
	}
}

func TestDivideOperation_Execute_InvalidOperands(t *testing.T) {
	operation := DivideOperation{}

	_, err := operation.Execute([]float64{10})

	if err != calculator.ErrInvalidOperands {
		t.Errorf("expected ErrInvalidOperands, got %v", err)
	}
}

func TestPowerOperation_Execute(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		expected float64
	}{
		{
			name:     "positive exponent",
			operands: []float64{2, 3},
			expected: 8,
		},
		{
			name:     "zero exponent",
			operands: []float64{10, 0},
			expected: 1,
		},
		{
			name:     "negative exponent",
			operands: []float64{2, -2},
			expected: 0.25,
		},
		{
			name:     "decimal base",
			operands: []float64{2.5, 2},
			expected: 6.25,
		},
	}

	operation := PowerOperation{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := operation.Execute(test.operands)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result != test.expected {
				t.Errorf("expected result %v, got %v", test.expected, result)
			}
		})
	}
}

func TestPowerOperation_Execute_InvalidOperands(t *testing.T) {
	tests := [][]float64{
		{2},
		{2, 3, 4},
	}

	operation := PowerOperation{}

	for _, operands := range tests {
		_, err := operation.Execute(operands)

		if err != calculator.ErrInvalidOperands {
			t.Errorf("expected ErrInvalidOperands, got %v", err)
		}
	}
}

func TestSquareRootOperation_Execute(t *testing.T) {
	tests := []struct {
		name     string
		operands []float64
		expected float64
	}{
		{
			name:     "perfect square",
			operands: []float64{25},
			expected: 5,
		},
		{
			name:     "zero",
			operands: []float64{0},
			expected: 0,
		},
		{
			name:     "decimal result",
			operands: []float64{2.25},
			expected: 1.5,
		},
	}

	operation := SquareRootOperation{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := operation.Execute(test.operands)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result != test.expected {
				t.Errorf("expected result %v, got %v", test.expected, result)
			}
		})
	}
}

func TestSquareRootOperation_Execute_NegativeNumber(t *testing.T) {
	operation := SquareRootOperation{}

	_, err := operation.Execute([]float64{-1})

	if err != calculator.ErrNegativeSquareRoot {
		t.Errorf("expected ErrNegativeSquareRoot, got %v", err)
	}
}

func TestSquareRootOperation_Execute_InvalidOperands(t *testing.T) {
	tests := [][]float64{
		{},
		{25, 4},
	}

	operation := SquareRootOperation{}

	for _, operands := range tests {
		_, err := operation.Execute(operands)

		if err != calculator.ErrInvalidOperands {
			t.Errorf("expected ErrInvalidOperands, got %v", err)
		}
	}
}
