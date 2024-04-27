# Calc builder
BuildСalcFunc builds the СalcFunc function. CalcFunc effectively calculates the value of the expression passed to BuildСalcFunc during construction. Expression can contain variables. CalcFunc obtains the values of variables during calculation using the getVar callback function.

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
    getVar := func(i int) float64 {return vars[i]}

    // (x0 * x1) / (x0 + x1)
    calc, err := calcbuilder.BuildСalcFunc("/ * x0 x1 + x0 x1", getVar) 
    if err != nil {
        // error handling
    }

    result := calc() // (70 * 30) / (70 + 30)

    fmt.Println(result) // 21
}
```

## Benchmark
[benchmark code](https://github.com/guamoko995/calcbuilder/blob/master/tests/calcbuilder_bench_test.go)

### Test data

#### Expression
```(x0 * x1) / (x0 + x1)```

#### Function for getting variable values
```go
// x0 = 70, x1 = 30
vars = []float64{70, 30}
getVar = func(i int) float64 { return vars[i] }
```

### Сompared cases
1. Used calc builder.
2. Used compilled calc function.
    ```go
    func compilled() float64 {
        return getVar(0) * getVar(1) / (getVar(0) + getVar(1))
    }
   ```
3. Used compilled calc function in var.
    ```go
	compilledInVar := func() float64 {
		return getVar(0) * getVar(1) / (getVar(0) + getVar(1))
	}
   ```
4. Used a [stack calculator](https://github.com/guamoko995/calcbuilder/blob/master/tests/calc_on_stack_test.go).

### Benchmark results
```
goos: linux
goarch: amd64
pkg: github.com/guamoko995/calcbuilder/tests
cpu: AMD Ryzen 5 5600H with Radeon Graphics         
Benchmark/calc_builder-12         	62482611	        19.07 ns/op	       0 B/op	       0 allocs/op
Benchmark/compilled-12            	179408836	         6.874 ns/op	       0 B/op	       0 allocs/op
Benchmark/compilled_in_var-12     	179180094	         6.717 ns/op	       0 B/op	       0 allocs/op
Benchmark/calc_on_stack-12        	12713480	        95.77 ns/op	      56 B/op	       7 allocs/op
PASS
ok  	github.com/guamoko995/calcbuilder/tests	7.313s
```

TODO Understand why there are allocations in runtime calculations of the benchmark.
