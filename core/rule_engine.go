package core

type RuleEngine struct {
	Rules  []*Rule
	Schema Schema
}

func (e *RuleEngine) AddRule(rule *Rule) error {
	if err := rule.Validate(&e.Schema); err != nil {
		return err
	}
	e.Rules = append(e.Rules, rule)
	return nil
}

func (e *RuleEngine) AddRules(rules []*Rule) error {
	for _, rule := range rules {
		if err := e.AddRule(rule); err != nil {
			return err
		}
	}
	return nil
}

func (e *RuleEngine) Validate() error {
	if err := e.Schema.Validate(); err != nil {
		return err
	}
	for _, rule := range e.Rules {
		if err := rule.Validate(&e.Schema); err != nil {
			return err
		}
	}
	return nil
}

func (e *RuleEngine) Run(fact map[string]string) *Event {
	for _, rule := range e.Rules {
		event := rule.Execute(fact)
		if event != nil {
			return event
		}
	}
	return nil
}
