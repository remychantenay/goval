package uuid

import (
	"fmt"
	"strings"
	"github.com/satori/go.uuid"
	"github.com/remychantenay/goval/generic"
	"github.com/remychantenay/goval/constant"
)


type UuidValidator struct {
	Required bool
}

func (v UuidValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	str := val.(string)
	l := len(str)
	expectedSize := 36

	if l == 0 {
		if v.Required {return false, fmt.Errorf("cannot be blank")}
	} else {
		if l != expectedSize {return false, fmt.Errorf("should be %v characters long", expectedSize)}
		_, err := uuid.FromString(str)
		if err != nil {return false, fmt.Errorf("invalid uuid")}
	}

	return true, nil
}

// BuildUuidValidator allows to build the validator for UUIDs
func BuildUuidValidator(args []string) generic.Validator {
	validator := UuidValidator{}
	count := len(args)
	for i := 0; i < count; i++ {
		fmt.Println(args[i])
		if strings.Contains(args[i], constant.ArgConstraintRequired) {
			fmt.Sscanf(args[i], constant.ArgConstraintRequired+"%t", &validator.Required)
		}
	}
	return validator
}

// BuildUuidValidator allows to build the validator for UUIDs, mainly used for unit tests
func BuildUuidValidatorWithFullTag(tag string) generic.Validator {
	args := strings.Split(tag, ",")
	return BuildUuidValidator(args[1:])
}