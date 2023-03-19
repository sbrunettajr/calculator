package calculator

import "errors"

var ErrorDivisionByZero = errors.New("it isn't possible divide by zero")

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Mul(a, b float64) float64 {
	return a * b
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrorDivisionByZero
	}
	return a / b, nil
}
