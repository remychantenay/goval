package string

import (
	"fmt"
	"strings"
	"sync"

	"github.com/remychantenay/goval/generic"
)

const (
	// argConstraintRequired defines whether or not the field is required.
	// In this context, required means not nil and not blank.
	argConstraintRequired = "required="

	// argConstraintMax defines the maximum number of characters allowed.
	argConstraintMax = "max="

	// argConstraintMin defines the minimum number of characters allowed.
	argConstraintMin = "min="
)

var pool = sync.Pool{
	New: func() interface{} { return &stringValidator{-1, -1, false} },
}

type stringValidator struct {
	min      int
	max      int
	required bool
}

// Validate a specific field
func (v *stringValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.required {
		return false, fmt.Errorf("cannot be nil")
	}

	l := len(val.(string))

	if l == 0 && v.required {
		return false, fmt.Errorf("cannot be blank")
	}

	if v.min != -1 && l < v.min {
		return false, fmt.Errorf("should be at least %d characters long", v.min)
	}

	if v.max != -1 && v.min != -1 && v.max >= v.min && l > v.max {
		return false, fmt.Errorf("should be less than %d characters long", v.max)
	}

	return true, nil
}

// NewValidator builds the validator for strings
func NewValidator(args []string) generic.Validator {
	validator := pool.Get().(*stringValidator)
	defer pool.Put(validator)
	if validator.max != -1 {
		validator.min, validator.max = -1, -1
		validator.required = false
	}

	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], argConstraintMax) {
			fmt.Sscanf(args[i], argConstraintMax+"%d", &validator.max)
		} else if strings.Contains(args[i], argConstraintMin) {
			fmt.Sscanf(args[i], argConstraintMin+"%d", &validator.min)
		} else if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &validator.required)
		}
	}
	return validator
}

// NewValidatorWithoutPool builds the validator for strings
// Only here for benchmark purposes
func NewValidatorWithoutPool(args []string) generic.Validator {
	validator := stringValidator{-1, -1, false}

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
