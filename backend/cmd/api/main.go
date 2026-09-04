package main

import (
	apiCalculator "calculator/backend/internal/api/calculator"
	domainCalculator "calculator/backend/internal/calculator"
	"calculator/backend/internal/calculator/operations"
	"log"
	"net/http"
)

func main() {
	operations := map[domainCalculator.OperationType]domainCalculator.Operation{
		domainCalculator.Add:        operations.AddOperation{},
		domainCalculator.Subtract:   operations.SubtractOperation{},
		domainCalculator.Multiply:   operations.MultiplyOperation{},
		domainCalculator.Divide:     operations.DivideOperation{},
		domainCalculator.Power:      operations.PowerOperation{},
		domainCalculator.SquareRoot: operations.SquareRootOperation{},
	}

	calc := domainCalculator.NewCalculator(operations)

	handler := apiCalculator.NewHandler(calc)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/calculator", handler.Calculate)

	log.Println("server listening on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
