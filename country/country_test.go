package country

import (
	"testing"
	"reflect"
	"github.com/remychantenay/goval/generic"
)

type TestCountry struct {
	Value 					string		`goval:"country_code,required=true,exclude=US,excludeEu=true"`
	ValueWithoutConstraint 	string		`goval:"country_code"`
}

func validate(testedStruct TestCountry, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	return BuildCountryCodeValidatorWithFullTag(generic.ExtractTag(structValue, index)).Validate(structValue.Field(index).Interface())
}

func TestCountryShort(t *testing.T) {
	_, actualResult := validate(TestCountry{"E", ""}, 0)
	var expectedResult = "should be 2 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryTooLong(t *testing.T) {
	_, actualResult := validate(TestCountry{"EEE", ""}, 0)
	var expectedResult = "should be 2 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryInvalid(t *testing.T) {
	_, actualResult := validate(TestCountry{"RR", ""}, 0)
	var expectedResult = "is an invalid country code"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryEmptyWithConstraint(t *testing.T) {
	_, actualResult := validate(TestCountry{"", ""}, 0)
	var expectedResult = "cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryExcludedWithConstraint(t *testing.T) {
	_, actualResult := validate(TestCountry{"US", ""}, 0)
	var expectedResult = "is excluded"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryOK(t *testing.T) {
	_, actualResult := validate(TestCountry{"CH", ""}, 0)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestCountryEmptyWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestCountry{"", ""}, 1)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}