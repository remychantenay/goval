package email

import (
	"github.com/remychantenay/goval/generic"
	"reflect"
	"strings"
	"testing"
)

// Email represents the struct under test
type Email struct {
	Value                  string `goval:"email,required=true,domain=google.com"`
	ValueWithoutConstraint string `goval:"email"`
}

func TestEmail_Err(t *testing.T) {
	tests := []struct {
		description      string
		expectedErrorMsg string
		with             Email
		index            int
	}{
		{
			description:      "Invalid",
			expectedErrorMsg: "is an invalid email address",
			with: Email{
				Value: "somethingnotvalid",
			},
			index: 0,
		},
		{
			description:      "Empty with constraint",
			expectedErrorMsg: "cannot be blank",
			with: Email{
				Value:                  "",
				ValueWithoutConstraint: "",
			},
			index: 0,
		},
		{
			description:      "Invalid domain",
			expectedErrorMsg: "is an invalid domain",
			with: Email{
				Value: "john.smith@facebook.com",
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

func TestEmail_Success(t *testing.T) {
	_, err := validate(Email{"john.smith@google.com", ""}, 0)

	if err != nil {
		t.Fatalf("No error expected but got %s", err.Error())
	}
}

// validate is a convenience function
func validate(testedStruct Email, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	tag := generic.ExtractTag(structValue, index)
	args := strings.Split(tag, ",")
	return NewValidator(args[1:]).Validate(structValue.Field(index).Interface())
}
