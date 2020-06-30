package string

import (
	"reflect"
	"strings"
	"testing"

	"github.com/remychantenay/goval/generic"
)

// String represents the struct under test
type String struct {
	Value                  string `goval:"string,min=10,max=15,required=true"` // With constraints
	Value2                 string `goval:"string,required=false"`
	Value3                 string `goval:"string,required=false"`
	Value4                 string `goval:"string,required=false"`
	ValueWithoutConstraint string `goval:"string"`
}

func TestString_Err(t *testing.T) {
	tests := []struct {
		description      string
		expectedErrorMsg string
		with             String
		index            int
	}{
		{
			description:      "Too Short",
			expectedErrorMsg: "should be at least 10 characters long",
			with: String{
				Value: "tooShort",
			},
			index: 0,
		},
		{
			description:      "Too long",
			expectedErrorMsg: "should be less than 15 characters long",
			with: String{
				Value: "wayTooooooooooooooooooLong",
			},
			index: 0,
		},
		{
			description:      "Empty",
			expectedErrorMsg: "cannot be blank",
			with: String{
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

func TestString_Success(t *testing.T) {
	tests := []struct {
		description string
		with        String
		index       int
	}{
		{
			description: "Success with constraint",
			with: String{
				Value: "ValidString",
			},
			index: 0,
		},
		{
			description: "Success without constraint",
			with: String{
				ValueWithoutConstraint: "",
			},
			index: 1,
		},
		{
			description: "Too short without constraint",
			with: String{
				ValueWithoutConstraint: "tooShort",
			},
			index: 1,
		},
		{
			description: "Too long without constraint",
			with: String{
				ValueWithoutConstraint: "wayTooooooooooooooooooLong",
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

func BenchmarkString_WithoutPool(b *testing.B) {
	b.StopTimer()

	tests := []struct {
		description string
		with        String
		index       int
	}{
		{
			description: "Success with constraint",
			with: String{
				Value:  "ValidString",
				Value2: "Another String",
				Value3: "And Another One",
				Value4: "One More Time",
			},
			index: 0,
		},
		{
			description: "Success without constraint",
			with: String{
				ValueWithoutConstraint: "",
				Value2:                 "Another String",
				Value3:                 "And Another One",
				Value4:                 "One More Time",
			},
			index: 1,
		},
		{
			description: "Too short without constraint",
			with: String{
				ValueWithoutConstraint: "tooShort",
				Value2:                 "Another String",
				Value3:                 "And Another One",
				Value4:                 "One More Time",
			},
			index: 1,
		},
		{
			description: "Too long without constraint",
			with: String{
				ValueWithoutConstraint: "wayTooooooooooooooooooLong",
				Value2:                 "Another String",
				Value3:                 "And Another One",
				Value4:                 "One More Time",
			},
			index: 1,
		},
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for _, test := range tests {
			structValue := reflect.ValueOf(test.with)
			tag := generic.ExtractTag(structValue, test.index)
			args := strings.Split(tag, ",")
			NewValidatorWithoutPool(args[1:]).Validate(structValue.Field(test.index).Interface())
		}
	}
	b.StopTimer()
}

func BenchmarkString_WithPool(b *testing.B) {
	b.StopTimer()
	tests := []struct {
		description string
		with        String
		index       int
	}{
		{
			description: "Success with constraint",
			with: String{
				Value:  "ValidString",
				Value2: "Another String",
				Value3: "And Another One",
				Value4: "One More Time",
			},
			index: 0,
		},
		{
			description: "Success without constraint",
			with: String{
				ValueWithoutConstraint: "",
				Value2:                 "Another String",
				Value3:                 "And Another One",
				Value4:                 "One More Time",
			},
			index: 1,
		},
		{
			description: "Too short without constraint",
			with: String{
				ValueWithoutConstraint: "tooShort",
				Value2:                 "Another String",
				Value3:                 "And Another One",
				Value4:                 "One More Time",
			},
			index: 1,
		},
		{
			description: "Too long without constraint",
			with: String{
				ValueWithoutConstraint: "wayTooooooooooooooooooLong",
				Value2:                 "Another String",
				Value3:                 "And Another One",
				Value4:                 "One More Time",
			},
			index: 1,
		},
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		for _, test := range tests {
			validate(test.with, test.index)
		}
	}
	b.StopTimer()
}

// validate is a convenience function
func validate(testedStruct String, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	tag := generic.ExtractTag(structValue, index)
	args := strings.Split(tag, ",")
	return NewValidator(args[1:]).Validate(structValue.Field(index).Interface())
}
