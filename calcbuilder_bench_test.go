package calcbuilder

import (
	"testing"
)

var (
	vars    []float64
	getVars func(i int) float64
)

func Benchmark(b *testing.B) {
	// x0 = 70, x1 = 30
	vars = []float64{70, 30}
	getVars = func(i int) float64 { return vars[i] }

	b.Run("calc builder", func(b *testing.B) {
		// (x0 * x1) / (x0 + x1)
		calc, err := BuildСalcFunc("/ * x0 x1 + x0 x1", getVars)
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

	b.Run("compiled", func(b *testing.B) {
		if compiled() != 21 {
			panic("wrong answer")
		}

		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			_ = compiled() // (70 * 30) / (70 + 30)
		}
		b.StopTimer()
	})

	b.Run("compiled in var", func(b *testing.B) {
		// (x0 * x1) / (x0 + x1)
		compilInVar := func() float64 {
			return getVars(0) * getVars(1) / (getVars(0) + getVars(1))
		}

		if compilInVar() != 21 {
			panic("wrong answer")
		}

		b.ResetTimer()
		for i := 1; i <= b.N; i++ {
			_ = compilInVar() // (70 * 30) / (70 + 30)
		}
		b.StopTimer()
	})

	b.Run("calc runtime", func(b *testing.B) {
		// (x0 * x1) / (x0 + x1)
		c, err := newCalcRuntime("x0 x1 * x0 x1 + /", getVars)
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
func compiled() float64 {
	return getVars(0) * getVars(1) / (getVars(0) + getVars(1))
}
