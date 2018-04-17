package string

import (
	"testing"
	"github.com/remychantenay/goval/generic"
	"reflect"
)

type TestString struct {
	Value 					string		`goval:"string,min=10,max=15,required=true"` // With constraints
	ValueWithoutConstraint 	string		`goval:"string"`
}

func validate(testedStruct TestString, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	return BuildStringValidatorWithFullTag(generic.ExtractTag(structValue, index)).Validate(structValue.Field(index).Interface())
}

func TestStringShortWithConstraint(t *testing.T) {
	_, actualResult := validate(TestString{"tooShort", ""}, 0)
	var expectedResult = "should be at least 10 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestStringTooLongWithConstraint(t *testing.T) {
	_, actualResult := validate(TestString{"wayTooooooooooooooooooLong", ""}, 0)
	var expectedResult = "should be less than 15 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestStringEmptyWithConstraint(t *testing.T) {
	_, actualResult := validate(TestString{"", ""}, 0)
	var expectedResult = "cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestStringOK(t *testing.T) {
	_, actualResult := validate(TestString{"ValidString", ""}, 0)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestStringShortWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestString{"", "tooShort"}, 1)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestStringTooLongWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestString{"", "wayTooooooooooooooooooLong"}, 1)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestStringEmptyWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestString{"", ""}, 1)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}