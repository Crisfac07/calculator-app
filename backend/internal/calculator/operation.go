package calculator

type Operation interface {
	Execute(operands []float64) (float64, error)
}
