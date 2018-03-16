package goval

import (
	"fmt"
	"reflect"
	"strings"
)

const nameTag = "goval"

type Validator interface {
	Validate(interface{}) (bool, error)
}

type GenericValidator struct {}

func (v GenericValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

func main() {

}

func ValidateStruct(s interface{}) []error {
	errs := []error{}

	structValue := reflect.ValueOf(s)

	for i := 0; i < structValue.NumField(); i++ {
		tag := structValue.Type().Field(i).Tag.Get(nameTag)

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidator(tag)

		valid, err := validator.Validate(structValue.Field(i).Interface())

		// Append error to results
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", structValue.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}

func getValidator(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case ARG_TYPE_NUMBER:
		return buildStringValidator(args[1:])
	case ARG_TYPE_STRING:
		return buildNumberValidator(args[1:])
	case ARG_TYPE_EMAIL:
		return buildEmailValidator(args[1:])
	}

	return GenericValidator{}
}
