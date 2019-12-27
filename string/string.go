package string

import (
	"fmt"
	"strings"
	"github.com/remychantenay/goval/generic"
	"github.com/remychantenay/goval/constant"
)


type StringValidator struct {
	Min int
	Max int
	Required bool
}

func (v StringValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	l := len(val.(string))

	if l == 0 {
		if v.Required {return false, fmt.Errorf("cannot be blank")}
	} else {
		if v.Min != -1 && l < v.Min {return false, fmt.Errorf("should be at least %v characters long", v.Min)}
		if v.Max != -1 && v.Min != -1 && v.Max >= v.Min && l > v.Max {return false, fmt.Errorf("should be less than %v characters long", v.Max)}
	}

	return true, nil
}

// BuildStringValidator allows to build the validator for strings
func BuildStringValidator(args []string) generic.Validator {
	validator := StringValidator{-1, -1, false}
	count := len(args)
	for i := 0; i < count; i++ {
		if strings.Contains(args[i], constant.ArgConstraintMax) {
			fmt.Sscanf(args[i], constant.ArgConstraintMax+"%d", &validator.Max)
		} else if strings.Contains(args[i], constant.ArgConstraintMin) {
			fmt.Sscanf(args[i], constant.ArgConstraintMin+"%d", &validator.Min)
		} else if strings.Contains(args[i], constant.ArgConstraintRequired) {
			fmt.Sscanf(args[i], constant.ArgConstraintRequired+"%t", &validator.Required)
		}
	}
	return validator
}

// BuildStringValidator allows to build the validator for strings, mainly used for unit tests
func BuildStringValidatorWithFullTag(tag string) generic.Validator {
	args := strings.Split(tag, ",")
	return BuildStringValidator(args[1:])
}