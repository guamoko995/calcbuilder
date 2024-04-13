package calcbuilder

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
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
func BuildСalcFunc(expression string) (calcFunc func(getVar func(i int) float64) float64, err error) {
	calcFunc, apendix, err := buildСalcFunc(expression)
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

func buildСalcFunc(expression string) (СalcFunc func(getVar func(i int) float64) float64, apendix string, err error) {
	if expression == "" {
		return nil, "", ErrUnexpectedEndOfExpression
	}
	token, apendix, _ := strings.Cut(expression, " ")
	switch token {
	// TO DO support more operators (including unary)
	case "+":
		var f1, f2 func(func(i int) float64) float64

		position := len(expression) - len(apendix) - 1
		f1, apendix, err = buildСalcFunc(apendix)
		if err != nil {
			var tg PositionErr
			if errors.As(err, &tg) {
				tg.Position += position
				err = tg
			}
			return nil, "", err
		}
		position = len(expression) - len(apendix) - 1
		f2, apendix, err = buildСalcFunc(apendix)
		if err != nil {
			var tg PositionErr
			if errors.As(err, &tg) {
				tg.Position += position
				err = tg
			}
			return nil, "", err
		}

		return func(getVar func(i int) float64) float64 {
			return f1(getVar) + f2(getVar)
		}, apendix, nil
	case "-":
		var f1, f2 func(func(i int) float64) float64

		position := len(expression) - len(apendix) - 1
		f1, apendix, err = buildСalcFunc(apendix)
		if err != nil {
			var tg PositionErr
			if errors.As(err, &tg) {
				tg.Position += position
				err = tg
			}
			return nil, "", err
		}
		position = len(expression) - len(apendix) - 1
		f2, apendix, err = buildСalcFunc(apendix)
		if err != nil {
			var tg PositionErr
			if errors.As(err, &tg) {
				tg.Position += position
				err = tg
			}
			return nil, "", err
		}

		return func(getVar func(i int) float64) float64 {
			return f1(getVar) - f2(getVar)
		}, apendix, nil
	case "*":
		var f1, f2 func(func(i int) float64) float64

		position := len(expression) - len(apendix) - 1
		f1, apendix, err = buildСalcFunc(apendix)
		if err != nil {
			var tg PositionErr
			if errors.As(err, &tg) {
				tg.Position += position
				err = tg
			}
			return nil, "", err
		}
		position = len(expression) - len(apendix) - 1
		f2, apendix, err = buildСalcFunc(apendix)
		if err != nil {
			var tg PositionErr
			if errors.As(err, &tg) {
				tg.Position += position
				err = tg
			}
			return nil, "", err
		}

		return func(getVar func(i int) float64) float64 {
			return f1(getVar) * f2(getVar)
		}, apendix, nil
	case "/":
		var f1, f2 func(func(i int) float64) float64

		position := len(expression) - len(apendix) - 1
		f1, apendix, err = buildСalcFunc(apendix)
		if err != nil {
			var tg PositionErr
			if errors.As(err, &tg) {
				tg.Position += position
				err = tg
			}
			return nil, "", err
		}
		position = len(expression) - len(apendix) - 1
		f2, apendix, err = buildСalcFunc(apendix)
		if err != nil {
			var tg PositionErr
			if errors.As(err, &tg) {
				tg.Position += position
				err = tg
			}
			return nil, "", err
		}

		return func(getVar func(i int) float64) float64 {
			return f1(getVar) / f2(getVar)
		}, apendix, nil
	default:
		if token[0] == 'x' { // var
			i64, err := strconv.ParseInt(token[1:], 10, 64)
			if err != nil {
				return nil, apendix, PositionErr{WrappedErr: ErrInvalidToken}
			}
			i := int(i64)
			return func(getVar func(i int) float64) float64 {
				return getVar(i)
			}, apendix, nil
		} else { // const
			c, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return nil, apendix, PositionErr{WrappedErr: ErrInvalidToken}
			}
			return func(getVar func(i int) float64) float64 {
				return c
			}, apendix, nil
		}
	}
}
