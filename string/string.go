package string

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

type stringValidator struct {
	Min      int
	Max      int
	Required bool
}

// Validate a specific field
func (v *stringValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required {
		return false, fmt.Errorf("cannot be nil")
	}

	l := len(val.(string))

	if l == 0 {
		if v.Required {
			return false, fmt.Errorf("cannot be blank")
		}
	} else {
		if v.Min != -1 && l < v.Min {
			return false, fmt.Errorf("should be at least %v characters long", v.Min)
		}
		if v.Max != -1 && v.Min != -1 && v.Max >= v.Min && l > v.Max {
			return false, fmt.Errorf("should be less than %v characters long", v.Max)
		}
	}

	return true, nil
}

// NewValidator builds the validator for strings
func NewValidator(args []string) generic.Validator {
	validator := stringValidator{-1, -1, false}
	count := len(args)
	for i := 0; i < count; i++ {
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
