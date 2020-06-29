package uuid

import (
	"fmt"
	"github.com/remychantenay/goval/generic"
	"github.com/satori/go.uuid"
	"strings"
)

const (
	// argConstraintRequired defines whether or not the field is required.
	// In this context, required means not nil and not blank.
	argConstraintRequired = "required="
)

type uuidValidator struct {
	required bool
}

// Validate a specific field
func (v *uuidValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.required {
		return false, fmt.Errorf("cannot be nil")
	}

	str := val.(string)
	l := len(str)
	expectedSize := 36 // UUID

	if l == 0 && v.required {
		return false, fmt.Errorf("cannot be blank")
	}

	if l != expectedSize {
		return false, fmt.Errorf("should be %d characters long", expectedSize)
	}
	_, err := uuid.FromString(str)
	if err != nil {
		return false, fmt.Errorf("invalid uuid")
	}

	return true, nil
}

// NewValidator builds and returns the validator for UUIDs.
func NewValidator(args []string) generic.Validator {
	validator := uuidValidator{}
	count := len(args)
	for i := 0; i < count; i++ {
		if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &validator.required)
		}
	}
	return &validator
}
