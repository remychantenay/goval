package country

import (
	"github.com/remychantenay/goval/generic"
	"reflect"
	"strings"
	"testing"
)

// Country represents the struct under test
type Country struct {
	Value                  string `goval:"country_code,required=true,exclude=US,excludeEU=true"`
	ValueWithoutConstraint string `goval:"country_code"`
}

func TestCountry_Err(t *testing.T) {
	tests := []struct {
		description      string
		expectedErrorMsg string
		with             Country
		index            int
	}{
		{
			description:      "Too short",
			expectedErrorMsg: "should be 2 characters long",
			with: Country{
				Value:                  "E",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Too long",
			expectedErrorMsg: "should be 2 characters long",
			with: Country{
				Value:                  "EEE",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Invalid",
			expectedErrorMsg: "is an invalid country code",
			with: Country{
				Value:                  "RR",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Empty with constraint",
			expectedErrorMsg: "cannot be blank",
			with: Country{
				Value:                  "",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Excluded country code",
			expectedErrorMsg: "is excluded",
			with: Country{
				Value:                  "US",
				ValueWithoutConstraint: "",
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

func TestCountry_Success(t *testing.T) {
	_, err := validate(Country{"CH", ""}, 0)

	if err != nil {
		t.Fatalf("No error expected but got %s", err.Error())
	}
}

// validate is a convenience function
func validate(testedStruct Country, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	tag := generic.ExtractTag(structValue, index)
	args := strings.Split(tag, ",")
	return NewValidator(args[1:]).Validate(structValue.Field(index).Interface())
}
