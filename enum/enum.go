package enum

import (
	"fmt"
	"strings"
	"github.com/remychantenay/goval/generic"
	"github.com/remychantenay/goval/constant"
)


type EnumValidator struct {
	Required bool
	Values string
}

func (v EnumValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	str := val.(string)
	l := len(str)

	if l == 0  && v.Required {return false, fmt.Errorf("cannot be blank")}

	b, err := inValues(str, v.Values)
	if !b { return b, err }

	return true, nil
}

func inValues(str string, valueList string) (bool, error) {
	if len(valueList) != 0 {
		valueArray := strings.Split(valueList, "|")
		valueArraySize := len(valueArray)
		for i := 0; i < valueArraySize; i++ {
			if str == valueArray[i] {return true, nil}
		}
	}

	return false, fmt.Errorf("is an invalid value: %s", str)
}

// BuildEnumValidator allows to build the validator for enums
func BuildEnumValidator(args []string) generic.Validator {
	validator := EnumValidator{false, ""}
	count := len(args)
	for i := 0; i < count; i++ {
		if strings.Contains(args[i], constant.ArgConstraintRequired) {
			fmt.Sscanf(args[i], constant.ArgConstraintRequired+"%t", &validator.Required)
		} else if strings.Contains(args[i], constant.ArgConstraintValues) {
			fmt.Sscanf(args[i], constant.ArgConstraintValues+"%s", &validator.Values)
		}
	}
	return validator
}

// BuildEnumValidator allows to build the validator for enums, mainly used for unit tests
func BuildEnumValidatorWithFullTag(tag string) generic.Validator {
	args := strings.Split(tag, ",")
	return BuildEnumValidator(args[1:])
}