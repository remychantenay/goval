package number

import (
	"fmt"
	"strings"
	"github.com/remychantenay/goval/generic"
	"github.com/remychantenay/goval/constant"
)

type NumberValidator struct {
	Min int
	Max int
	Required bool
}

func (v NumberValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	num := val.(int)

	if v.Min != -1 && num < v.Min {return false, fmt.Errorf("should be greater than %v", v.Min)}
	if v.Max != -1 && v.Max >= v.Min && num > v.Max {return false, fmt.Errorf("should be less than %v", v.Max)}

	return true, nil
}

// BuildNumberValidator allows to build the validator for numbers
func BuildNumberValidator(args []string) generic.Validator {
	validator := NumberValidator{-1, -1, false}
	count := len(args)-1
	for i := 0; i <= count; i++ {
		fmt.Println(args[i])
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

// BuildNumberValidator allows to build the validator for numbers, mainly used for unit tests
func BuildNumberValidatorWithFullTag(tag string) generic.Validator {
	args := strings.Split(tag, ",")
	return BuildNumberValidator(args[1:])
}