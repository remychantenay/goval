package email

import (
	"testing"
	"reflect"
	"github.com/remychantenay/goval/generic"
)

type TestEmail struct {
	Value 					string		`goval:"email,required=true,domain=google.com"`
	ValueWithoutConstraint 	string		`goval:"email"`
}

func validate(testedStruct TestEmail, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	return BuildEmailValidatorWithFullTag(generic.ExtractTag(structValue, index)).Validate(structValue.Field(index).Interface())
}

func TestEmailEmptyWithConstraint(t *testing.T) {
	_, actualResult := validate(TestEmail{"", ""}, 0)
	var expectedResult = "cannot be blank"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestEmailValid(t *testing.T) {
	_, actualResult := validate(TestEmail{"john.smith@google.com", ""}, 0)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestEmailInvalid(t *testing.T) {
	_, actualResult := validate(TestEmail{"somethingnotvalid", ""}, 0)
	var expectedResult = "is an invalid email address"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestEmailEmptyWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestEmail{"", ""}, 1)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestEmailInvalidDomainWithConstraint(t *testing.T) {
	_, actualResult := validate(TestEmail{"john.smith@facebook.com", ""}, 0)
	var expectedResult = "is an invalid domain"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}