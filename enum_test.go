package goval

import (
	"testing"
)

type TestEnum struct {
	Value 	string		`goval:"enum,required=true,values=SOMETHING|SOMETHING_ELSE"` // With constraints
}

type TestEnum2 struct {
	Value 	string		`goval:"enum"` // Without constraints
}

func TestEnumEmptyWithConstraint(t *testing.T) {
	actualResult := ValidateStruct(TestEnum{""})[0]
	var expectedResult = "Value cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestEnumValid(t *testing.T) {
	actualResult := ValidateStruct(TestEnum{"SOMETHING"})

	if len(actualResult) != 0 {
		t.Fatalf("No error expected but got %v error(s)", len(actualResult))
	}
}

func TestEnumInvalid(t *testing.T) {
	actualResult := ValidateStruct(TestEnum{"SOMETHING_INVALID"})[0]
	var expectedResult = "Value is an invalid value: SOMETHING_INVALID"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}