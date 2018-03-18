package goval

import (
	"testing"
)

type TestEmail struct {
	Value 	string		`goval:"email,required=true,domain=google.com"` // With constraints
}

type TestEmail2 struct {
	Value 	string		`goval:"email"` // Without constraints
}

func TestEmailEmptyWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestEmail{""})[0]
	var expectedResult = "Value cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestEmailValid(t *testing.T) {
	actualResult := ValidateStruct(TestEmail{"john.smith@google.com"})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestEmailInvalid(t *testing.T) {
	actualResult := ValidateStruct(TestEmail{"somethingnotvalid"})[0]
	var expectedResult = "Value is an invalid email address"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestEmailEmptyWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestEmail2{""})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestEmailInvalidDomainWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestEmail{"john.smith@facebook.com"})[0]
	var expectedResult = "Value is an invalid domain"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}