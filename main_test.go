package goval

import (
	"testing"
)

type TestCountryCode struct {
	Value 					string		`goval:"country_code,required=true,exclude=US,excludeEu=true"`
	ValueWithoutConstraint 	string		`goval:"country_code"`
}

func TestCountryShort(t *testing.T) {
	actualResult := ValidateStruct(TestCountryCode{"E", ""})[0]
	var expectedResult = "Value should be 2 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryTooLong(t *testing.T) {
	actualResult := ValidateStruct(TestCountryCode{"EEEE", ""})[0]
	var expectedResult = "Value should be 2 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryInvalid(t *testing.T) {
	actualResult := ValidateStruct(TestCountryCode{"RR", ""})[0]
	var expectedResult = "Value is an invalid country code"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryEmptyWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestCountryCode{"", ""})[0]
	var expectedResult = "Value cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryExcludedWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestCountryCode{"US", ""})[0]
	var expectedResult = "Value is excluded"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryOK(t *testing.T) {
	actualResult := ValidateStruct(TestCountryCode{"CH", ""})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestCountryEmptyWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestCountryCode{"CH", ""})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}