package uuid

import (
	"github.com/remychantenay/goval/generic"
	"reflect"
	"strings"
	"testing"
)

// UUID represents the struct under test
type UUID struct {
	Value                  string `goval:"uuid,required=true"`
	ValueWithoutConstraint string `goval:"uuid"`
}

func TestUUID_Err(t *testing.T) {
	tests := []struct {
		description      string
		expectedErrorMsg string
		with             UUID
		index            int
	}{
		{
			description:      "Too Short with constraint",
			expectedErrorMsg: "should be 36 characters long",
			with: UUID{
				Value: "47bd",
			},
			index: 0,
		},
		{
			description:      "Too long with constraint",
			expectedErrorMsg: "should be 36 characters long",
			with: UUID{
				Value: "f025b018-a0cb-47bd-97ce-f460f20e3b255555555",
			},
			index: 0,
		},
		{
			description:      "Too Short without constraint",
			expectedErrorMsg: "should be 36 characters long",
			with: UUID{
				ValueWithoutConstraint: "47bd",
			},
			index: 1,
		},
		{
			description:      "Too long without constraint",
			expectedErrorMsg: "should be 36 characters long",
			with: UUID{
				ValueWithoutConstraint: "f025b018-a0cb-47bd-97ce-f460f20e3b255555555",
			},
			index: 1,
		},
		{
			description:      "Empty",
			expectedErrorMsg: "cannot be blank",
			with: UUID{
				Value: "",
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

func TestUUID_Success(t *testing.T) {
	tests := []struct {
		description string
		with        UUID
		index       int
	}{
		{
			description: "Success with constraint",
			with: UUID{
				Value: "f025b018-a0cb-47bd-97ce-f460f20e3b25",
			},
			index: 0,
		},
		{
			description: "Empty without constraint",
			with: UUID{
				ValueWithoutConstraint: "",
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
func validate(testedStruct UUID, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	tag := generic.ExtractTag(structValue, index)
	args := strings.Split(tag, ",")
	return NewValidator(args[1:]).Validate(structValue.Field(index).Interface())
}
