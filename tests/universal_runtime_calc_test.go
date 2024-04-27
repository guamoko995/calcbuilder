package calcbuilder_test

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type calcRuntime struct {
	stack  []any
	frml   []any
	getVar func(int) float64
}

func newCalcRuntime(expression string, getVar func(int) float64) (*calcRuntime, error) {
	terms := strings.Split(expression, " ")
	c := &calcRuntime{
		stack:  make([]any, len(terms)),
		frml:   make([]any, 0, len(terms)),
		getVar: getVar,
	}
	for _, term := range terms {
		switch term {
		case "+", "-", "*", "/":
			c.frml = append(c.frml, term)
		default:
			if term[0] == 'x' {
				v, err := strconv.ParseInt(term[1:], 10, 64)
				if err != nil {
					return nil, errors.Wrap(err, "failed parse var")
				}
				c.frml = append(c.frml, int(v))
			} else {
				v, err := strconv.ParseFloat(term, 64)
				if err != nil {
					return nil, errors.Wrap(err, "failed parse var")
				}
				c.frml = append(c.frml, v)
			}
		}
	}
	return c, nil
}

// TODO Understand why there are allocations.
func (c *calcRuntime) calc() float64 {
	lenStack := 0
	for _, terml := range c.frml {
		switch s := (terml).(type) {
		case float64:
			c.stack[lenStack] = s
			lenStack++
		case int:
			v := c.getVar(s)
			c.stack[lenStack] = v
			lenStack++
		case string:
			switch s {
			case "+", "-", "*", "/":
				if lenStack == 0 {
					panic("formula line is not valid")
				}
				arg2 := c.stack[lenStack-1].(float64)
				lenStack--

				if lenStack == 0 {
					panic("formula line is not valid")
				}
				arg1 := c.stack[lenStack-1].(float64)
				lenStack--

				switch s {
				case "+":
					c.stack[lenStack] = arg1 + arg2
				case "-":
					c.stack[lenStack] = arg1 - arg2
				case "*":
					c.stack[lenStack] = arg1 * arg2
				case "/":
					c.stack[lenStack] = arg1 / arg2
				default:
					panic("formula line is not valid")
				}
				lenStack++
			default:
				panic("formula line is not valid")
			}
		default:
			panic("formula line is not valid")
		}
	}
	if lenStack != 1 {
		panic("formula line is not valid")
	}
	return c.stack[0].(float64)
}
