package number

import (
	"fmt"
	"github.com/remychantenay/goval/generic"
	"strings"
)

const (
	// argConstraintRequired
	argConstraintRequired = "required="

	// argConstraintMax
	argConstraintMax = "max="

	// argConstraintMin
	argConstraintMin = "min="
)

type numberValidator struct {
	Min      int64
	Max      int64
	Required bool
}

// Validate a specific field
func (v *numberValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required {
		return false, fmt.Errorf("cannot be nil")
	}

	num := val.(int64)

	if v.Min != -1 && num < v.Min {
		return false, fmt.Errorf("should be greater than %v", v.Min)
	}
	if v.Max != -1 && v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("should be less than %v", v.Max)
	}

	return true, nil
}

// NewValidator allows to build the validator for numbers
func NewValidator(args []string) generic.Validator {
	validator := numberValidator{-1, -1, false}
	count := len(args) - 1
	for i := 0; i <= count; i++ {
		if strings.Contains(args[i], argConstraintMax) {
			fmt.Sscanf(args[i], argConstraintMax+"%d", &validator.Max)
		} else if strings.Contains(args[i], argConstraintMin) {
			fmt.Sscanf(args[i], argConstraintMin+"%d", &validator.Min)
		} else if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &validator.Required)
		}
	}
	return &validator
}
