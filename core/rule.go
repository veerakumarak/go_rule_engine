package core

type Rule struct {
	Condition Condition
	Event     Event
}

func (r *Rule) Validate(schema *Schema) error {
	return r.Condition.Validate(schema)
}

func (r *Rule) Execute(fact map[string]string) *Event {
	if r.Condition.Evaluate(fact) {
		return &r.Event
	}
	return nil
}
