package goval

import (
	"testing"
)

type TestString struct {
	Value 	string		`goval:"string,min=10,max=15,required=true"` // With constraints
}

type TestString2 struct {
	Value 	string		`goval:"string"` // Without constraints
}

func TestStringShortWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestString{"tooShort"})[0]
	var expectedResult = "Value should be at least 10 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestStringTooLongWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestString{"wayTooooooooooooooooooLong"})[0]
	var expectedResult = "Value should be less than 15 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestStringEmptyWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestString{""})[0]
	var expectedResult = "Value cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestStringOK(t *testing.T) {
	actualResult := ValidateStruct(TestString{"ValidString"})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestStringShortWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestString2{"tooShort"})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestStringTooLongWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestString2{"wayTooooooooooooooooooLong"})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestStringEmptyWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestString2{""})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}