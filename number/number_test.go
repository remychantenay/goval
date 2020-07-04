package number

import (
	"reflect"
	"strings"
	"testing"

	"github.com/remychantenay/goval/generic"
)

// Number represents the struct under test
type Number struct {
	Value                  int64 `goval:"number,min=10,max=15,required=true"`
	ValueWithoutConstraint int64 `goval:"number"`
}

func TestNumber_Err(t *testing.T) {
	tests := []struct {
		description      string
		expectedErrorMsg string
		with             Number
		index            int
	}{
		{
			description:      "Too small with constraint",
			expectedErrorMsg: "should be greater than 10",
			with: Number{
				Value: 9,
			},
			index: 0,
		},
		{
			description:      "Too large with constraint",
			expectedErrorMsg: "should be less than 15",
			with: Number{
				Value: 16,
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

func TestNumber_Success(t *testing.T) {
	tests := []struct {
		description string
		with        Number
		index       int
	}{
		{
			description: "Success with or without constraint",
			with: Number{
				Value: 13,
			},
			index: 0,
		},
		{
			description: "Too small without constraint",
			with: Number{
				Value: 9,
			},
			index: 1,
		},
		{
			description: "Too large without constraint",
			with: Number{
				Value: 16,
			},
			index: 1,
		},
	}

	for _, test := range tests {
		_, err := validate(test.with, test.index)
		if err != nil {
			t.Fatalf("%s -> was not expecting error but got %s", test.description, err.Error())
		}
	}
}

// validate is a convenience function
func validate(testedStruct Number, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	tag := generic.ExtractTag(structValue, index)
	args := strings.Split(tag, ",")
	return NewValidator(args[1:]).Validate(structValue.Field(index).Interface())
}
