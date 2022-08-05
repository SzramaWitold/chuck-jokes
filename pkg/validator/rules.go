package validator

import (
	"fmt"
	"strconv"
)

var v = map[string]func(i string) error{
	"required": required,
	"uint":     isUint,
}

func required(i string) error {
	if i == "" {
		return fmt.Errorf("field required")
	}
	return nil
}

func isUint(i string) error {
	inputInt, convErr := strconv.Atoi(i)

	if convErr != nil {
		return fmt.Errorf("field should be of numeric type")
	}

	if inputInt < 0 {
		return fmt.Errorf("field should be positive number")
	}

	return nil
}
