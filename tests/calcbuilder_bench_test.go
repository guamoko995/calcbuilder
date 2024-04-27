package calcbuilder_test

import (
	"testing"

	"github.com/guamoko995/calcbuilder"
)

var (
	vars   []float64
	getVar func(i int) float64
)

func Benchmark(b *testing.B) {
	// x0 = 70, x1 = 30
	vars = []float64{70, 30}
	getVar = func(i int) float64 { return vars[i] }

	b.Run("calc builder", func(b *testing.B) {
		// (x0 * x1) / (x0 + x1)
		calc, err := calcbuilder.BuildСalcFunc("/ * x0 x1 + x0 x1", getVar)
		if err != nil {
			panic(err)
		}

		if calc() != 21 {
			panic("wrong answer")
		}

		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			_ = calc() // (70 * 30) / (70 + 30)
		}
		b.StopTimer()
	})

	b.Run("compilled", func(b *testing.B) {
		if compilled() != 21 {
			panic("wrong answer")
		}

		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			_ = compilled() // (70 * 30) / (70 + 30)
		}
		b.StopTimer()
	})

	b.Run("compilled in var", func(b *testing.B) {
		// (x0 * x1) / (x0 + x1)
		compilledInVar := func() float64 {
			return getVar(0) * getVar(1) / (getVar(0) + getVar(1))
		}

		if compilledInVar() != 21 {
			panic("wrong answer")
		}

		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			_ = compilledInVar() // (70 * 30) / (70 + 30)
		}
		b.StopTimer()
	})

	b.Run("calc on stack", func(b *testing.B) {
		// (x0 * x1) / (x0 + x1)
		c, err := newCalcOnStack("x0 x1 * x0 x1 + /", getVar)
		if err != nil {
			panic(err)
		}

		if c.calc() != 21 {
			panic("wrong answer")
		}

		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			_ = c.calc() // (70 * 30) / (70 + 30)
		}
		b.StopTimer()
	})
}

// (x0 * x1) / (x0 + x1)
func compilled() float64 {
	return getVar(0) * getVar(1) / (getVar(0) + getVar(1))
}
