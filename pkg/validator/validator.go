package validator

import (
	"fmt"
	"reflect"
	"strings"
)

type Validator struct {
	rules map[string]func(n, i string) error
}

func NewValidator() *Validator {
	return &Validator{
		rules: rulesMap,
	}
}

func (val *Validator) getRules(name, input string, rules ...string) []error {
	var errors []error
	for _, rule := range rules {
		err := val.rules[rule](name, input)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func (val *Validator) Validate(request interface{}, inputs map[string]string) []error {
	t := reflect.TypeOf(request)
	var errors []error
	if t.Kind() != reflect.Struct {
		return append(errors, fmt.Errorf("wrong request type"))
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		value := inputs[f.Name]
		rules := strings.Split(f.Tag.Get("validation"), ",")
		errors := val.getRules(f.Name, value, rules...)
		if errors != nil {
			return errors
		}
	}

	return nil
}
