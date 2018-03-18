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

// ValidateStruct allows to validate any struct containing the nameTag
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

// getValidator will return the appropriate validator
func getValidator(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case ArgTypeNumber:
		return buildNumberValidator(args[1:])
	case ArgTypeString:
		return buildStringValidator(args[1:])
	case ArgTypeEmail:
		return buildEmailValidator(args[1:])
	case ArgTypeUuid:
		return buildUuidValidator(args[1:])
	case ArgTypeCountryCode:
		return buildCountryCodeValidator(args[1:])
	case ArgTypeCurrency:
		return buildCurrencyValidator(args[1:])
	case ArgTypeEnum:
		return buildEnumValidator(args[1:])
	}

	return GenericValidator{}
}
