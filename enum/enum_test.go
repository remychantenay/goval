package enum

import (
	"testing"
	"reflect"
	"github.com/remychantenay/goval/generic"
)

type TestEnum struct {
	Value 					string		`goval:"enum,required=true,values=SOMETHING|SOMETHING_ELSE"`
	ValueWithoutConstraint 	string		`goval:"enum"`
}

func validate(testedStruct TestEnum, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	return BuildEnumValidatorWithFullTag(generic.ExtractTag(structValue, index)).Validate(structValue.Field(index).Interface())
}

func TestEnumEmptyWithConstraint(t *testing.T) {
	_, actualResult := validate(TestEnum{"", ""}, 0)
	var expectedResult = "cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestEnumValid(t *testing.T) {
	_, actualResult := validate(TestEnum{"SOMETHING", ""}, 0)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestEnumInvalid(t *testing.T) {
	_, actualResult := validate(TestEnum{"SOMETHING_INVALID", ""}, 0)
	var expectedResult = "is an invalid value: SOMETHING_INVALID"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}