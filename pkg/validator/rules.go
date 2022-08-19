package validator

import (
	"fmt"
	"strconv"
)

var rulesMap = map[string]func(n, i string) error{
	"required": required,
	"uint":     isUint,
}

func required(n, i string) error {
	if i == "" {
		return fmt.Errorf("field '%v' required", n)
	}
	return nil
}

func isUint(n, i string) error {
	inputInt, convErr := strconv.Atoi(i)

	if convErr != nil {
		return fmt.Errorf("field '%v' should be of numeric type", n)
	}

	if inputInt <= 0 {
		return fmt.Errorf("field '%v' should be positive number", n)
	}

	return nil
}
