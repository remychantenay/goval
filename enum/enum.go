package enum

import (
	"fmt"
	"github.com/remychantenay/goval/generic"
	"strings"
)

const (
	// argConstraintRequired
	argConstraintRequired = "required="

	// argConstraintDomain (e.g. @google.com)
	argConstraintDomain = "domain="

	// argConstraintValues: Exclusion parameters must be separated by a pipe (e.g. exclude=EUR|GBP)
	argConstraintValues = "values="
)

type enumValidator struct {
	Required bool
	Values   string
}

// Validate validate a specific field
func (v *enumValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required {
		return false, fmt.Errorf("cannot be nil")
	}

	str := val.(string)
	l := len(str)

	if l == 0 && v.Required {
		return false, fmt.Errorf("cannot be blank")
	}

	b, err := inValues(str, v.Values)
	if !b {
		return b, err
	}

	return true, nil
}

func inValues(str string, valueList string) (bool, error) {
	if len(valueList) != 0 {
		valueArray := strings.Split(valueList, "|")
		valueArraySize := len(valueArray)
		for i := 0; i < valueArraySize; i++ {
			if str == valueArray[i] {
				return true, nil
			}
		}
	}

	return false, fmt.Errorf("is an invalid value: %s", str)
}

// NewValidator build and returns the validator for enums
func NewValidator(args []string) generic.Validator {
	validator := enumValidator{false, ""}
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &validator.Required)
		} else if strings.Contains(args[i], argConstraintValues) {
			fmt.Sscanf(args[i], argConstraintValues+"%s", &validator.Values)
		}
	}
	return &validator
}
