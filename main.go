package goval

import (
	"fmt"
	"reflect"
	"strings"
	"github.com/remychantenay/goval/country"
	"github.com/remychantenay/goval/currency"
	"github.com/remychantenay/goval/generic"
	"github.com/remychantenay/goval/constant"
	"github.com/remychantenay/goval/email"
	"github.com/remychantenay/goval/uuid"
	"github.com/remychantenay/goval/enum"
	"github.com/remychantenay/goval/number"
	str "github.com/remychantenay/goval/string"
)


// ValidateStruct allows to validate any struct containing the nameTag
func ValidateStruct(s interface{}) []error {
	errs := []error{}

	structValue := reflect.ValueOf(s)

	for i := 0; i < structValue.NumField(); i++ {
		tag := generic.ExtractTag(structValue, i)

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
func getValidator(tag string) generic.Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case constant.ArgTypeNumber:
		return number.BuildNumberValidator(args[1:])
	case constant.ArgTypeString:
		return str.BuildStringValidator(args[1:])
	case constant.ArgTypeEmail:
		return email.BuildEmailValidator(args[1:])
	case constant.ArgTypeUuid:
		return uuid.BuildUuidValidator(args[1:])
	case constant.ArgTypeCountryCode:
		return country.BuildCountryCodeValidator(args[1:])
	case constant.ArgTypeCurrency:
		return currency.BuildCurrencyValidator(args[1:])
	case constant.ArgTypeEnum:
		return enum.BuildEnumValidator(args[1:])
	}

	return generic.GenericValidator{}
}
