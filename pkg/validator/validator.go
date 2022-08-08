package validator

type Validator struct {
	Rules map[string]func(n, i string) error
}

func NewValidator() *Validator {
	return &Validator{
		Rules: rulesMap,
	}
}

func (val *Validator) Validate(name, input string, rules ...string) []error {
	var errors []error
	for _, rule := range rules {
		err := val.Rules[rule](name, input)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
