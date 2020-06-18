package currency

import (
	"github.com/remychantenay/goval/generic"
	"reflect"
	"strings"
	"testing"
)

// Currency represents the struct under test
type Currency struct {
	Value                  string `goval:"currency,required=true,exclude=EUR|GBP"`
	ValueWithoutConstraint string `goval:"currency"`
}

func TestCurrency_Err(t *testing.T) {
	tests := []struct {
		description      string
		expectedErrorMsg string
		with             Currency
		index            int
	}{
		{
			description:      "Too short",
			expectedErrorMsg: "should be 3 characters long",
			with: Currency{
				Value:                  "EU",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Too long",
			expectedErrorMsg: "should be 3 characters long",
			with: Currency{
				Value:                  "EURR",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Invalid",
			expectedErrorMsg: "is an invalid currency",
			with: Currency{
				Value:                  "GGG",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Empty with constraint",
			expectedErrorMsg: "cannot be blank",
			with: Currency{
				Value:                  "",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Excluded currency",
			expectedErrorMsg: "is excluded",
			with: Currency{
				Value:                  "GBP",
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

func TestCurrency_Success(t *testing.T) {
	_, err := validate(Currency{"USD", ""}, 0)

	if err != nil {
		t.Fatalf("No error expected but got %s", err.Error())
	}
}

// validate is a convenience function
func validate(testedStruct Currency, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	tag := generic.ExtractTag(structValue, index)
	args := strings.Split(tag, ",")
	return NewValidator(args[1:]).Validate(structValue.Field(index).Interface())
}
