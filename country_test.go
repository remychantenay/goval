package goval

import (
	"testing"
)

type TestCountry struct {
	Value 	string		`goval:"country_code,required=true,exclude=US,excludeEu=true"` // With constraints
}

type TestCountry2 struct {
	Value 	string		`goval:"country_code"` // Without constraints
}

func TestCountryShort(t *testing.T) {
	actualResult := ValidateStruct(TestCountry{"E"})[0]
	var expectedResult = "Value should be 2 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryTooLong(t *testing.T) {
	actualResult := ValidateStruct(TestCountry{"EEE"})[0]
	var expectedResult = "Value should be 2 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryInvalid(t *testing.T) {
	actualResult := ValidateStruct(TestCountry{"RR"})[0]
	var expectedResult = "Value is an invalid country code"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryEmptyWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestCountry{""})[0]
	var expectedResult = "Value cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryExcludedWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestCountry{"US"})[0]
	var expectedResult = "Value is excluded"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestCountryOK(t *testing.T) {
	actualResult := ValidateStruct(TestCountry{"CH"})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestCountryEmptyWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestCountry2{""})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}