package core

type Schema struct {
	//Properties
	Properties map[string]Property
}

type Property struct {
	Type     string
	Priority int
}

func (s *Schema) Validate() error {
	return nil
}
