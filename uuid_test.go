package goval

import (
	"testing"
)

type TestUuid struct {
	Value 	string		`goval:"uuid,required=true"` // With constraints
}

type TestUuid2 struct {
	Value 	string		`goval:"uuid"` // Without constraints
}

func TestUuidShortWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestUuid{"f025b018-a0cb-47bd-97ce-"})[0]
	var expectedResult = "Value should be 36 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidTooLongWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestUuid{"f025b018-a0cb-47bd-97ce-f460f20e3b255555555"})[0]
	var expectedResult = "Value should be 36 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidEmptyWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestUuid{""})[0]
	var expectedResult = "Value cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidOK(t *testing.T) {
	actualResult := ValidateStruct(TestUuid{"f025b018-a0cb-47bd-97ce-f460f20e3b25"})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestUuidShortWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestUuid2{"f025b018-a0cb-47bd-97ce-"})[0]
	var expectedResult = "Value should be 36 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidTooLongWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestUuid2{"f025b018-a0cb-47bd-97ce-f460f20e3b25555555"})[0]
	var expectedResult = "Value should be 36 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidEmptyWithoutConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestUuid2{""})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}