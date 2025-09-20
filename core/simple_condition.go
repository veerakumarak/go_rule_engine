package core

import (
	"errors"
	"fmt"
	"reflect"
)

func (c *Condition) IsValidSimple(schema *Schema) error {
	if c.Operator == "" {
		return errors.New("operator should not be empty for a condition")
	}

	if c.Value == "" {
		return errors.New("value should not be empty for a condition")
	}

	if c.Fact == "" {
		return errors.New("fact should not be empty for a condition")
	}

	if c.All != nil || c.Any != nil {
		return errors.New("all and any are not supported for a condition")
	}

	property, ok := schema.Properties[c.Fact]
	if !ok {
		return errors.New(fmt.Sprintf("fact %s doesnt exists in schema", c.Fact))
	}

	if property.Type != reflect.TypeOf(c.Value).String() {
		return errors.New(fmt.Sprintf("expected value type %s, but given %s", property.Type, reflect.TypeOf(c.Value).String()))
	}

	return IsValidOperator(c.Operator)
}

func (c *Condition) EvaluateSimple(fact map[string]string) bool {
	if _, ok := fact[c.Fact]; !ok {
		return false
	}
	v, _ := fact[c.Fact]
	switch c.Operator {
	case Eq:
		return c.Value == v
	case Neq:
		return c.Value != v
	case Lt:
		return c.Value < v
	case Gt:
		return c.Value > v
	case Le:
		return c.Value <= v
	case Ge:
		return c.Value >= v
	}
	return false
}
