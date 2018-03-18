package goval

import (
	"testing"
)

type TestCurrency struct {
	Value 	string		`goval:"currency,required=true,exclude=EUR|GBP"` // With constraints
}

type TestCurrency2 struct {
	Value 	string		`goval:"currency"` // Without constraints
}

func TestCurrencyShort(t *testing.T) {
	actualResult := ValidateStruct(TestCurrency{"EU"})[0]
	var expectedResult = "Value should be 3 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyTooLong(t *testing.T) {
	actualResult := ValidateStruct(TestCurrency{"EURR"})[0]
	var expectedResult = "Value should be 3 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyInvalid(t *testing.T) {
	actualResult := ValidateStruct(TestCurrency{"GGG"})[0]
	var expectedResult = "Value is an invalid currency"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyEmptyWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestCurrency{""})[0]
	var expectedResult = "Value cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyExcludedWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestCurrency{"GBP"})[0]
	var expectedResult = "Value is excluded"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCurrencyOK(t *testing.T) {
	actualResult := ValidateStruct(TestCurrency{"USD"})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestCurrencyEmptyWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestCurrency2{""})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}