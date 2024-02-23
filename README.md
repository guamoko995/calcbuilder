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

TO DO add more operators and functions

## Usage

```go
package main

import (
	"calcbuilder"
	"fmt"
)

func main() {
    // (x0 * x1) / (x0 + x1)
	calc, err := calcbuilder.BuildСalcFunc("/ * x0 x1 + x0 x1 ") 
	if err != nil {
		// error handling
	}

	// x0 = 5, x1 = 10
	vars := []float64{70, 30}
	getVars := func(i int) float64 {return vars[i]}

	result := calc(getVars) // (70 * 30) / (70 + 30)

	fmt.Println(result) // 21
}
```

## Benchmarks
TO DO add bench results