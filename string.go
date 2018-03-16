package goval

import (
	"fmt"
	"strings"
)


type StringValidator struct {
	Min int
	Max int
	Required bool
}

func (v StringValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	l := len(val.(string))

	if l == 0  && v.Required {return false, fmt.Errorf("cannot be blank")}
	if l < v.Min {return false, fmt.Errorf("should be at least %v characters long", v.Min)}
	if v.Max >= v.Min && l > v.Max {return false, fmt.Errorf("should be less than %v characters long", v.Max)}
	return true, nil
}

func buildStringValidator(args []string) Validator {
	validator := StringValidator{}
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