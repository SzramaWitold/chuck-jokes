package validator

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type Validator struct {
	db *gorm.DB
}

func NewValidator(db *gorm.DB) *Validator {
	return &Validator{
		db: db,
	}
}

func (val *Validator) getRules(name, input string, rules ...string) []error {
	var errors []error
	for _, rule := range rules {
		r := strings.Split(rule, ":")
		if r[0] == rule {
			result := reflect.ValueOf(val).MethodByName(rule).Call([]reflect.Value{reflect.ValueOf(name), reflect.ValueOf(input), reflect.ValueOf("")})

			err := result[0].Interface()

			if err != nil {
				errors = append(errors, err.(error))
			}
		} else {
			result := reflect.ValueOf(val).MethodByName(r[0]).Call([]reflect.Value{reflect.ValueOf(name), reflect.ValueOf(input), reflect.ValueOf(r[1])})
			err := result[0].Interface()
			if err != nil {
				errors = append(errors, err.(error))
			}
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
		if rules[0] == "" {
			return nil
		}
		errors := val.getRules(f.Name, value, rules...)
		if errors != nil {
			return errors
		}
	}

	return nil
}
