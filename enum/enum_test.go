package enum

import (
	"reflect"
	"strings"
	"testing"

	"github.com/remychantenay/goval/generic"
)

// Enum represents the struct under test
type Enum struct {
	Value                  string `goval:"enum,required=true,values=SOMETHING|SOMETHING_ELSE"`
	ValueWithoutConstraint string `goval:"enum"`
}

func TestEnum_Err(t *testing.T) {
	tests := []struct {
		description      string
		expectedErrorMsg string
		with             Enum
		index            int
	}{
		{
			description:      "Empty with constraint",
			expectedErrorMsg: "cannot be blank",
			with: Enum{
				Value:                  "",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Invalid value",
			expectedErrorMsg: "is an invalid value: SOMETHING_INVALID",
			with: Enum{
				Value: "SOMETHING_INVALID",
			},
			index: 0,
		},
	}

	for _, test := range tests {
		_, err := validate(test.with, test.index)
		if err == nil {
			t.Fatalf("%s -> was expecting %s", test.description, test.expectedErrorMsg)
		}

		if err.Error() != test.expectedErrorMsg {
			t.Fatalf("%s -> Got %s but expected %s ", test.description, err.Error(), test.expectedErrorMsg)
		}
	}
}

func TestEnum_Success(t *testing.T) {
	_, err := validate(Enum{"SOMETHING", ""}, 0)

	if err != nil {
		t.Fatalf("No error expected but got %s", err.Error())
	}
}

// validate is a convenience function
func validate(testedStruct Enum, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	tag := generic.ExtractTag(structValue, index)
	args := strings.Split(tag, ",")
	return NewValidator(args[1:]).Validate(structValue.Field(index).Interface())
}
