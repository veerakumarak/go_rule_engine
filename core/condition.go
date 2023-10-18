package core

type Condition struct {
	Operator Operator
	Fact     string
	Value    string
	Any      []*Condition
	All      []*Condition
}

func (c *Condition) Validate(schema *Schema) error {
	if c.IsComposite() {
		return c.IsValidComposite(schema)
	}
	return c.IsValidSimple(schema)
}

func (c *Condition) Evaluate(fact map[string]string) bool {
	if c.Any != nil {
		return c.EvaluateAny(fact)
	} else if c.All != nil {
		return c.EvaluateAll(fact)
	} else {
		return c.EvaluateSimple(fact)
	}
}
