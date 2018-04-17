package number

import (
	"testing"
	"reflect"
	"github.com/remychantenay/goval/generic"
)

type TestNumber struct {
	Value 					int		`goval:"number,min=10,max=15,required=true"`
	ValueWithoutConstraint 	int		`goval:"number"`
}

func validate(testedStruct TestNumber, index int) (bool, error) {
	structValue := reflect.ValueOf(testedStruct)
	return BuildNumberValidatorWithFullTag(generic.ExtractTag(structValue, index)).Validate(structValue.Field(index).Interface())
}

func TestNumberTooSmallWithConstraint(t *testing.T) {
	_, actualResult := validate(TestNumber{9, 0}, 0)
	var expectedResult = "should be greater than 10"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}


func TestNumberTooLargeWithConstraint(t *testing.T) {
	_, actualResult := validate(TestNumber{16, 0}, 0)
	var expectedResult = "should be less than 15"

	if actualResult.Error() != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestNumberOK(t *testing.T) {
	_, actualResult := validate(TestNumber{13, 0}, 0)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestNumberTooSmallWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestNumber{0, 9}, 1)

	if actualResult != nil {
		t.Fatalf("No error expected but got %s", actualResult.Error())
	}
}

func TestNumberTooLongWithoutConstraint(t *testing.T) {
	_, actualResult := validate(TestNumber{0, 16}, 1)

	if actualResult != nil {
		t.Fatalf("No error expected but got %v error(s)", actualResult.Error())
	}
}
