package calculator

import (
	"calculator/backend/internal/calculator"
	"encoding/json"
	"net/http"
)

type Calculator interface {
	Calculate(
		operationType calculator.OperationType,
		operands []float64,
	) (float64, error)
}

type Handler struct {
	calculator Calculator
}

func NewHandler(calculator Calculator) *Handler {
	return &Handler{
		calculator: calculator,
	}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	var request CalculateRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"INVALID_REQUEST",
			"invalid request body",
		)
		return
	}

	result, err := h.calculator.Calculate(
		calculator.OperationType(request.Operation),
		request.Operands,
	)

	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			errorCode(err),
			err.Error(),
		)
		return
	}

	response := CalculateResponse{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
