package calculator

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	result := Add(1, 2)

	if result != 3 {
		t.Fatal() // Study
	}
}

func TestSub(t *testing.T) {
	t.Parallel()

	result := Sub(1, 2)

	if result != -1 {
		t.Fatal() // Study
	}
}

func TestMul(t *testing.T) {
	t.Parallel()

	result := Mul(3, 2)

	if result != 6 {
		t.Fatal()
	}
}

func TestDiv(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc   string
		a      float64
		b      float64
		result float64
		err    error
	}{
		{
			desc:   "ShouldReturnSuccess",
			a:      10,
			b:      5,
			result: 2,
		},
		{
			desc: "ShouldReturnErrorDivisionByZero",
			a:    1,
			b:    0,
			err:  ErrorDivisionByZero,
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			result, err := Div(tC.a, tC.b)

			if tC.err != nil {
				if !errors.Is(err, ErrorDivisionByZero) {
					t.Fatal() // Study
				}
			}

			if result != tC.result {
				t.Fatal() // Study
			}
		})
	}
}
