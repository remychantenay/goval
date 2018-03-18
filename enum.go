package goval

import (
	"fmt"
	"strings"
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

func buildEnumValidator(args []string) Validator {
	validator := EnumValidator{false, ""}
	count := len(args)
	for i := 0; i < count; i++ {
		fmt.Println(args[i])
		if strings.Contains(args[i], ArgConstraintRequired) {
			fmt.Sscanf(args[i], ArgConstraintRequired+"%t", &validator.Required)
		} else if strings.Contains(args[i], ArgConstraintValues) {
			fmt.Sscanf(args[i], ArgConstraintValues+"%s", &validator.Values)
		}
	}
	return validator
}
