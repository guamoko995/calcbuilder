package calcbuilder

import (
	"fmt"
	"strconv"
	"strings"
)

// BuildСalcFunc builds the СalcFunc function. CalcFunc effectively calculates
// the value of the expression passed to BuildСalcFunc during construction.
// expression can contain variables. CalcFunc obtains the values of variables during
// calculation using the getVars callback function.
//
// Expression format.
//   - Еhe expression is used in direct Polish notation.
//   - All operators, variables and constants in the expression are separated from each other
//     by one space.
//   - Constant syntax is defined by Go syntax for floating point literals.
//   - The variable is represented by the symbol 'x' (code 120) and the variable index
//     immediately following it. It is this index that will be passed to the getVals
//     callback to obtain the value of this variable during calculation.
//
// Valid operators: "+", "-" (binary), "*", "/"
func BuildСalcFunc(expression string) (СalcFunc func(getVar func(i int) float64) float64, err error) {
	// TO DO informative error handling without using panic and recovery
	defer func() {
		if recover() != nil {
			err = fmt.Errorf("expression not valid")
		}
	}()
	f, apendix := buildСalcFunc(expression)
	if len(apendix) > 0 {
		err = fmt.Errorf("unexpected continuation of expression")
	}
	return f, nil
}

func buildСalcFunc(expression string) (СalcFunc func(getVar func(i int) float64) float64, apendix string) {
	s := strings.SplitN(expression, " ", 2)
	o := s[0]
	if len(s) > 1 {
		expression = s[1]
	}
	switch o {
	// TO DO support more operators (including unary)
	case "+":
		var f1, f2 func(func(i int) float64) float64

		f1, expression = buildСalcFunc(expression)
		f2, expression = buildСalcFunc(expression)

		return func(getVar func(i int) float64) float64 {
			return f1(getVar) + f2(getVar)
		}, expression
	case "-":
		var f1, f2 func(func(i int) float64) float64

		f1, expression = buildСalcFunc(expression)
		f2, expression = buildСalcFunc(expression)

		return func(getVar func(i int) float64) float64 {
			return f1(getVar) - f2(getVar)
		}, expression
	case "*":
		var f1, f2 func(func(i int) float64) float64

		f1, expression = buildСalcFunc(expression)
		f2, expression = buildСalcFunc(expression)

		return func(getVar func(i int) float64) float64 {
			return f1(getVar) * f2(getVar)
		}, expression
	case "/":
		var f1, f2 func(func(i int) float64) float64

		f1, expression = buildСalcFunc(expression)
		f2, expression = buildСalcFunc(expression)

		return func(getVar func(i int) float64) float64 {
			return f1(getVar) / f2(getVar)
		}, expression
	default:
		if o[0] == 'x' { // var
			i64, _ := strconv.ParseInt(o[1:], 10, 64)
			i := int(i64)
			return func(getVar func(i int) float64) float64 {
				return getVar(i)
			}, expression
		} else { // const
			c, _ := strconv.ParseFloat(o, 64)

			return func(getVar func(i int) float64) float64 {
				return c
			}, expression
		}
	}
}
