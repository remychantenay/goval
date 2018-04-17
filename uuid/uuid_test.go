package uuid

import (
	"testing"
	"reflect"
	"github.com/remychantenay/goval/generic"
)

type TestUuid struct {
	Value 					string		`goval:"uuid,required=true"`
	ValueWithoutConstraint 	string		`goval:"uuid"`
}

func validate(testedStruct TestUuid, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	return BuildUuidValidatorWithFullTag(generic.ExtractTag(structValue, index)).Validate(structValue.Field(index).Interface())
}

func TestUuidShortWithConstraint(t *testing.T) {
	_, actualResult := validate(TestUuid{"47bd", ""}, 0)
	var expectedResult = "should be 36 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidTooLongWithConstraint(t *testing.T) {
	_, actualResult := validate(TestUuid{"f025b018-a0cb-47bd-97ce-f460f20e3b255555555", ""}, 0)
	var expectedResult = "should be 36 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidEmptyWithConstraint(t *testing.T) {
	_, actualResult := validate(TestUuid{"", ""}, 0)
	var expectedResult = "cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidOK(t *testing.T) {
	_, actualResult := validate(TestUuid{"f025b018-a0cb-47bd-97ce-f460f20e3b25", ""}, 0)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestUuidShortWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestUuid{"", "f025b018-a0cb-47bd-97ce-"}, 1)
	var expectedResult = "should be 36 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidTooLongWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestUuid{"", "f025b018-a0cb-47bd-97ce-f460f20e3b25555555"}, 1)
	var expectedResult = "should be 36 characters long"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestUuidEmptyWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestUuid{"", ""}, 1)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}