package goval

import (
	"fmt"
	"strings"
)

type NumberValidator struct {
	Min int
	Max int
	Required bool
}

func (v NumberValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	num := val.(int)

	if num < v.Min {return false, fmt.Errorf("should be greater than %v", v.Min)}
	if v.Max >= v.Min && num > v.Max {return false, fmt.Errorf("should be less than %v", v.Max)}

	return true, nil
}

func buildNumberValidator(args []string) Validator {
	validator := NumberValidator{}
	count := len(args)
	for i := 0; i <= count; i++ {
		fmt.Println(args[i])
		if strings.Contains(args[i], ARG_CONSTRAINT_MAX) {
			fmt.Sscanf(args[i],ARG_CONSTRAINT_MAX+"%d", &validator.Max)
		} else if strings.Contains(args[i], ARG_CONSTRAINT_MIN) {
			fmt.Sscanf(args[i],ARG_CONSTRAINT_MIN+"%d", &validator.Min)
		} else if strings.Contains(args[i], ARG_CONSTRAINT_REQUIRED) {
			fmt.Sscanf(args[i],ARG_CONSTRAINT_REQUIRED+"%t", &validator.Required)
		}
	}
	return validator
}