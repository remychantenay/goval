package uuid

import (
	"fmt"
	"github.com/remychantenay/goval/generic"
	"github.com/satori/go.uuid"
	"strings"
)

const (
	// argConstraintRequired
	argConstraintRequired = "required="
)

type uuidValidator struct {
	Required bool
}

// Validate a specific field
func (v *uuidValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required {
		return false, fmt.Errorf("cannot be nil")
	}

	str := val.(string)
	l := len(str)
	expectedSize := 36

	if l == 0 {
		if v.Required {
			return false, fmt.Errorf("cannot be blank")
		}
	} else {
		if l != expectedSize {
			return false, fmt.Errorf("should be %v characters long", expectedSize)
		}
		_, err := uuid.FromString(str)
		if err != nil {
			return false, fmt.Errorf("invalid uuid")
		}
	}

	return true, nil
}

// NewValidator allows to build the validator for UUIDs
func NewValidator(args []string) generic.Validator {
	validator := uuidValidator{}
	count := len(args)
	for i := 0; i < count; i++ {
		if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &validator.Required)
		}
	}
	return &validator
}
