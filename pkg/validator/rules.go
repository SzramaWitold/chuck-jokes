package validator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func (val *Validator) Required(n, i, o string) error {
	if i == "" {
		return fmt.Errorf("field '%v' required", n)
	}
	return nil
}

func (val *Validator) IsUint(n, i, o string) error {
	if i == "" {
		return nil
	}
	inputInt, convErr := strconv.Atoi(i)

	if convErr != nil {
		return fmt.Errorf("field '%v' should be of numeric type", n)
	}

	if inputInt <= 0 {
		return fmt.Errorf("field '%v' should be positive number", n)
	}

	return nil
}

func (val *Validator) Date(n, i, o string) error {
	if i == "" {
		return nil
	}
	layout := "2022-02-01"
	_, err := time.Parse(layout, i)

	if err != nil {
		return fmt.Errorf("field '%v' should have format like: '%v'", n, layout)
	}

	return nil
}

func (val *Validator) Unique(n, i, o string) error {
	log.Println(o, n, i)
	var result struct {
		Found bool
	}

	query := fmt.Sprintf("SELECT EXISTS(SELECT * FROM %v WHERE %v = ?) AS found", o, strings.ToLower(n))
	val.db.Raw(query, i).Scan(&result)

	if result.Found == true {
		return fmt.Errorf("%v already taken", n)
	}
	return fmt.Errorf("twoja stara")
}
