package goval

import (
	"fmt"
	"strings"
	"github.com/satori/go.uuid"
)


type UuidValidator struct {
	Required bool
}

func (v UuidValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	str := val.(string)
	l := len(str)
	expectedSize := 36

	if l == 0  && v.Required {return false, fmt.Errorf("cannot be blank")}
	if l != expectedSize {return false, fmt.Errorf("should be %v characters long", expectedSize)}
	_, err := uuid.FromString(str)
	if err != nil {return false, fmt.Errorf("invalid uuid")}

	return true, nil
}

func buildUuidValidator(args []string) Validator {
	validator := UuidValidator{}
	count := len(args)
	for i := 0; i <= count; i++ {
		fmt.Println(args[i])
		if strings.Contains(args[i], ArgConstraintRequired) {
			fmt.Sscanf(args[i], ArgConstraintRequired+"%t", &validator.Required)
		}
	}
	return validator
}
