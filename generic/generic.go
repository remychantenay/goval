package generic

import (
	"fmt"
	"reflect"
	"strings"
)

// NameTag to add to the struct fields to be validated
const NameTag = "goval"

// Validator interface
type Validator interface {
	Validate(val interface{}) (bool, error)
}

type GenericValidator struct{}

// Validate will validates a given field.
func (v GenericValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

// ValueExcluded : returns true if the given value is excluded, false otherwise
func ValueExcluded(str string, excludeList string) (bool, error) {
	if len(excludeList) != 0 {

		// If one value (e.g. "GBP")
		if !strings.Contains(excludeList, "|") {
			if str == excludeList {
				return false, fmt.Errorf("is excluded")
			}
		} else { // Else (e.g. "GBP|EUR")
			excludeArray := strings.Split(excludeList, "|")
			excludeArraySize := len(excludeArray)
			for i := 0; i < excludeArraySize; i++ {
				if str == excludeArray[i] {
					return false, fmt.Errorf("is excluded")
				}
			}
		}
	}

	return true, nil
}

// ExtractTag allows to extract the tag of a given Struct field index
func ExtractTag(value reflect.Value, index int) string {
	return value.Type().Field(index).Tag.Get(NameTag)
}
