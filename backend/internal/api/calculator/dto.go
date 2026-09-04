package calculator

type CalculateRequest struct {
	Operation string    `json:"operation"`
	Operands  []float64 `json:"operands"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}
