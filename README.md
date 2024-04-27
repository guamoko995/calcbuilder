# Calc builder
BuildСalcFunc builds the СalcFunc function. CalcFunc effectively calculates the value of the expression passed to BuildСalcFunc during construction. Expression can contain variables. CalcFunc obtains the values of variables during calculation using the getVars callback function.

## Expression format
- The expression is a [normal Polish notation](https://en.wikipedia.org/wiki/Polish_notation).
- All operators, variables and constants in the expression are separated from each other by one space.
- Constant syntax is defined by Go syntax for floating point literals.
- The variable is represented by the symbol 'x' (code 120) and the variable index immediately following it. It is this index that will be passed to the getVals callback to obtain the value of this variable during calculation.

## Valid operators
- ```+``` addition
- ```-``` subtraction
- ```*``` multiplication
- ```/``` division

TODO add more operators and functions

## Example

```go
package main

import (
	"fmt"

	"github.com/guamoko995/calcbuilder"
)

func main() {
    // x0 = 70, x1 = 30
    vars := []float64{70, 30}
    getVars := func(i int) float64 {return vars[i]}

    // (x0 * x1) / (x0 + x1)
    calc, err := calcbuilder.BuildСalcFunc("/ * x0 x1 + x0 x1", getVars) 
    if err != nil {
        // error handling
    }

    result := calc() // (70 * 30) / (70 + 30)

    fmt.Println(result) // 21
}
```

## Benchmark
[benchmark code](https://github.com/guamoko995/calcbuilder/blob/master/calcbuilder_bench_test.go)

### Test expression
```(x0 * x1) / (x0 + x1)```

### Сompared cases
1. Used calc builder.
2. Used compilled calc function.
    ```go
    func compiled() float64 {
        return getVars(0) * getVars(1) / (getVars(0) + getVars(1))
    }
   ```
3. Used compilled calc function in var.
    ```go
	compilInVar := func() float64 {
		return getVars(0) * getVars(1) / (getVars(0) + getVars(1))
	}
   ```
4. Used a universal [runtime calculator](https://github.com/guamoko995/calcbuilder/blob/master/universal_runtime_calc.go).

### Benchmark results:
```
goos: linux
goarch: amd64
pkg: github.com/guamoko995/calcbuilder
cpu: AMD Ryzen 5 5600H with Radeon Graphics         
Benchmark/calc_builder-12         	61060794	        19.77 ns/op	       0 B/op	       0 allocs/op
Benchmark/compiled-12             	216661424	         5.583 ns/op	       0 B/op	       0 allocs/op
Benchmark/compiled_in_var-12      	214518717	         5.562 ns/op	       0 B/op	       0 allocs/op
Benchmark/calc_runtime-12         	12651817	        94.51 ns/op	      56 B/op	       7 allocs/op
PASS
ok  	github.com/guamoko995/calcbuilder	7.015s
```

TODO Understand why there are allocations in runtime calculations of the benchmark.
