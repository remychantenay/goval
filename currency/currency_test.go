package currency

import (
	"testing"
	"reflect"
	"github.com/remychantenay/goval/generic"
)

type TestCurrency struct {
	Value 					string		`goval:"currency,required=true,exclude=EUR|GBP"`
	ValueWithoutConstraint 	string		`goval:"currency"`
}

func validate(testedStruct TestCurrency, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	return BuildCurrencyValidatorWithFullTag(generic.ExtractTag(structValue, index)).Validate(structValue.Field(index).Interface())
}

func TestCurrencyShort(t *testing.T) {
	_, actualResult := validate(TestCurrency{"EU", ""}, 0)
	var expectedResult = "should be 3 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyTooLong(t *testing.T) {
	_, actualResult := validate(TestCurrency{"EURR", ""}, 0)
	var expectedResult = "should be 3 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyInvalid(t *testing.T) {
	_, actualResult := validate(TestCurrency{"GGG", ""}, 0)
	var expectedResult = "is an invalid currency"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyEmptyWithConstraint(t *testing.T) {
	_, actualResult := validate(TestCurrency{"", ""}, 0)
	var expectedResult = "cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyExcludedWithConstraint(t *testing.T) {
	_, actualResult := validate(TestCurrency{"GBP", ""}, 0)
	var expectedResult = "is excluded"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyOK(t *testing.T) {
	_, actualResult := validate(TestCurrency{"USD", ""}, 0)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestCurrencyEmptyWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestCurrency{"EU", ""}, 1)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}