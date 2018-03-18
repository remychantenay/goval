package goval

import (
	"testing"
)

type TestNumber struct {
	Value 	int		`goval:"number,min=10,max=15,required=true"` // With constraints
}

type TestNumber2 struct {
	Value 	int		`goval:"number"` // Without constraints
}

func TestNumberTooSmallWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestNumber{9})[0]
	var expectedResult = "Value should be greater than 10"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestNumberTooLargeWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestNumber{16})[0]
	var expectedResult = "Value should be less than 15"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestNumberOK(t *testing.T) {
	actualResult := ValidateStruct(TestNumber{13})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestNumberTooSmallWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestNumber2{9})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestNumberTooLongWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestNumber2{16})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}