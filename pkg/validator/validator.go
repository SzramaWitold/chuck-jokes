package validator

type Validator struct {
	Input  string
	Errors []error
	Valid  bool
}

func NewValidator() *Validator {
	return &Validator{}
}

func (val *Validator) Validate(input string, rules ...string) {
	for _, r := range rules {
		val.Valid = true
		err := v[r](input)
		if err != nil {
			val.Errors = append(val.Errors, err)
			val.Valid = false
		}
	}
}
