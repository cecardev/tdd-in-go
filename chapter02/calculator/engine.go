package calculator

type Engine struct{}

func (e *Engine) Add(a, b float64) float64 {
	return a + b
}
