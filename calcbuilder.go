package calcbuilder

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// BuildСalcFunc builds the СalcFunc function. CalcFunc effectively calculates
// the value of the expression passed to BuildСalcFunc during construction.
// expression can contain variables. CalcFunc obtains the values of variables during
// calculation using the getVar callback function.
//
// Expression format.
//   - Еhe expression is used in direct Polish notation.
//   - All operators, variables and constants in the expression are separated from each other
//     by one space.
//   - Constant syntax is defined by Go syntax for floating point literals.
//   - The variable is represented by the symbol 'x' (code 120) and the variable index
//     immediately following it. It is this index that will be passed to the getVals
//     callback to obtain the value of this variable during calculation.
func BuildСalcFunc(expression string, getVar func(i int) float64) (CalcFunc func() float64, err error) {
	CalcFunc, apendix, err := buildСalcFunc(expression, getVar)
	if err != nil {
		return nil, err
	}
	if len(apendix) > 0 {
		return nil, PositionErr{
			WrappedErr: ErrUnexpectedContinuationOfExpression,
			Position:   len(expression) - len(apendix),
		}
	}

	return
}

// Valid operators
var binaryOperators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 {
		return x + y
	},
	"-": func(x, y float64) float64 {
		return x - y
	},
	"*": func(x, y float64) float64 {
		return x * y
	},
	"/": func(x, y float64) float64 {
		return x / y
	},
}

func buildСalcFunc(expression string, getVar func(i int) float64) (calcFunc func() float64, apendix string, err error) {
	if expression == "" {
		return nil, "", ErrUnexpectedEndOfExpression
	}
	term, apendix, _ := strings.Cut(expression, " ")

	// binary operator
	if operator, exist := binaryOperators[term]; exist {
		var operands [2]func() float64
		position := len(expression) - len(apendix) - 1

		for i := range operands {
			operands[i], apendix, err = buildСalcFunc(apendix, getVar)
			if err != nil {
				var tg PositionErr
				if errors.As(err, &tg) {
					tg.Position += position
					err = tg
				}
				return nil, "", err
			}
		}
		return func() float64 {
			return operator(operands[0](), operands[1]())
		}, apendix, nil
	}

	// var
	if term[0] == 'x' {
		i64, err := strconv.ParseInt(term[1:], 10, 64)
		if err != nil {
			return nil, apendix, PositionErr{WrappedErr: ErrInvalidTerm}
		}
		i := int(i64)
		return func() float64 {
			return getVar(i)
		}, apendix, nil
	}

	// const
	c, err := strconv.ParseFloat(term, 64)
	if err != nil {
		return nil, apendix, PositionErr{WrappedErr: ErrInvalidTerm}
	}
	return func() float64 {
		return c
	}, apendix, nil
}
