package calculator

import (
	"calculator/backend/internal/calculator"
	"calculator/backend/internal/calculator/operations"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakeCalculator struct{}

func (fakeCalculator) Calculate(
	operationType calculator.OperationType,
	operands []float64,
) (float64, error) {
	return 15, nil
}

func TestHandler_Calculate(t *testing.T) {
	handler := NewHandler(fakeCalculator{})

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/calculator",
		strings.NewReader(`{
			"operation": "add",
			"operands": [10, 5]
		}`),
	)

	recorder := httptest.NewRecorder()

	handler.Calculate(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", recorder.Code)
	}

	var response CalculateResponse

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Result != 15 {
		t.Errorf("expected result 15, got %v", response.Result)
	}
}

func TestHandler_InvalidRequestBody(t *testing.T) {
	handler := NewHandler(fakeCalculator{})

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/calculator",
		strings.NewReader(`{"operation":`),
	)

	recorder := httptest.NewRecorder()

	handler.Calculate(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", recorder.Code)
	}

	var response ErrorResponse

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Error.Code != "INVALID_REQUEST" {
		t.Errorf("expected error code INVALID_REQUEST, got %s", response.Error.Code)
	}
}

func TestHandler_DivisionByZero(t *testing.T) {
	calculator := calculator.NewCalculator(map[calculator.OperationType]calculator.Operation{
		calculator.Divide: operations.DivideOperation{},
	})

	handler := NewHandler(calculator)

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/calculator",
		strings.NewReader(`{
			"operation": "divide",
			"operands": [10, 0]
		}`),
	)

	recorder := httptest.NewRecorder()

	handler.Calculate(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", recorder.Code)
	}

	var response ErrorResponse

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Error.Code != "DIVISION_BY_ZERO" {
		t.Errorf(
			"expected error code DIVISION_BY_ZERO, got %s",
			response.Error.Code,
		)
	}
}

func newCalculator() *calculator.Calculator {
	return calculator.NewCalculator(map[calculator.OperationType]calculator.Operation{
		calculator.Add:    operations.AddOperation{},
		calculator.Divide: operations.DivideOperation{},
	})
}

func TestHandler_InvalidOperations(t *testing.T) {
	tests := []struct {
		name         string
		operation    string
		operands     []float64
		expectedCode string
	}{
		{
			name:         "invalid operation",
			operation:    "invalid",
			operands:     []float64{10, 5},
			expectedCode: "INVALID_OPERATION",
		},
		{
			name:         "invalid operands",
			operation:    "add",
			operands:     []float64{10},
			expectedCode: "INVALID_OPERANDS",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := NewHandler(newCalculator())

			requestBody := `{
				"operation": "` + test.operation + `",
				"operands": [10]
			}`

			request := httptest.NewRequest(
				http.MethodPost,
				"/api/v1/calculator",
				strings.NewReader(requestBody),
			)

			recorder := httptest.NewRecorder()

			handler.Calculate(recorder, request)

			if recorder.Code != http.StatusBadRequest {
				t.Errorf("expected status 400, got %d", recorder.Code)
			}

			var response ErrorResponse

			if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if response.Error.Code != test.expectedCode {
				t.Errorf(
					"expected error code %s, got %s",
					test.expectedCode,
					response.Error.Code,
				)
			}
		})
	}
}
