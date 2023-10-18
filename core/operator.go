package core

import (
	"errors"
	"fmt"
)

type Operator string

const (
	Eq  Operator = "eq"
	Neq Operator = "neq"
	Lt  Operator = "lt"
	Le  Operator = "le"
	Gt  Operator = "gt"
	Ge  Operator = "ge"
)

var validOperators = []Operator{Eq, Neq, Lt, Le, Gt, Ge}

func IsValidOperator(operator Operator) error {
	for _, op := range validOperators {
		if op == operator {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid operator %s", operator))
}
