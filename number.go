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
		if strings.Contains(args[i], ArgConstraintMax) {
			fmt.Sscanf(args[i], ArgConstraintMax+"%d", &validator.Max)
		} else if strings.Contains(args[i], ArgConstraintMin) {
			fmt.Sscanf(args[i], ArgConstraintMin+"%d", &validator.Min)
		} else if strings.Contains(args[i], ArgConstraintRequired) {
			fmt.Sscanf(args[i], ArgConstraintRequired+"%t", &validator.Required)
		}
	}
	return validator
}