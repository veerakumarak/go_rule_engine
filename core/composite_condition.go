package core

import "errors"

func (c *Condition) IsComposite() bool {
	return c.Any != nil || c.All != nil
}

func (c *Condition) IsValidComposite(schema *Schema) error {
	if c.Operator != "" {
		return errors.New("invalid operator for composite condition")
	}

	if c.Value != "" {
		return errors.New("invalid value for composite condition")
	}

	if c.Fact != "" {
		return errors.New("invalid fact for composite condition")
	}

	if c.All != nil && c.Any != nil {
		return errors.New("all and any are not supported simultaneously for a single condition")
	}

	if c.All != nil {
		return ValidateMulti(c.All, schema)
	}
	if c.Any != nil {
		return ValidateMulti(c.Any, schema)
	}

	return errors.New("internal server error")
}

func ValidateMulti(conditions []*Condition, schema *Schema) error {
	for _, condition := range conditions {
		if err := condition.Validate(schema); err != nil {
			return err
		}
	}
	return nil
}

func (c *Condition) EvaluateAll(fact map[string]string) bool {
	for _, cond := range c.All {
		if !cond.Evaluate(fact) {
			return false
		}
	}
	return true
}

func (c *Condition) EvaluateAny(fact map[string]string) bool {
	for _, cond := range c.Any {
		if cond.Evaluate(fact) {
			return true
		}
	}
	return false
}
