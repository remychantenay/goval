package number

import (
	"fmt"
	"github.com/remychantenay/goval/generic"
	"strings"
)

const (
	// argConstraintRequired defines whether or not the field is required.
	// In this context, required means not nil and not blank.
	argConstraintRequired = "required="

	// argConstraintMax defines the maximum value allowed.
	argConstraintMax = "max="

	// argConstraintMax defines the minimum value allowed.
	argConstraintMin = "min="
)

type numberValidator struct {
	min      int64
	max      int64
	required bool
}

// Validate a specific field
func (v *numberValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.required {
		return false, fmt.Errorf("cannot be nil")
	}

	num := val.(int64)

	if v.min != -1 && num < v.min {
		return false, fmt.Errorf("should be greater than %d", v.min)
	}
	if v.max != -1 && v.max >= v.min && num > v.max {
		return false, fmt.Errorf("should be less than %d", v.max)
	}

	return true, nil
}

// NewValidator allows to build the validator for numbers
func NewValidator(args []string) generic.Validator {
	validator := numberValidator{-1, -1, false}

	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], argConstraintMax) {
			fmt.Sscanf(args[i], argConstraintMax+"%d", &validator.max)
		} else if strings.Contains(args[i], argConstraintMin) {
			fmt.Sscanf(args[i], argConstraintMin+"%d", &validator.min)
		} else if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &validator.required)
		}
	}
	return &validator
}
