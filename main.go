package goval

import (
	"fmt"
	"github.com/remychantenay/goval/country"
	"github.com/remychantenay/goval/currency"
	"github.com/remychantenay/goval/email"
	"github.com/remychantenay/goval/enum"
	"github.com/remychantenay/goval/generic"
	"github.com/remychantenay/goval/number"
	str "github.com/remychantenay/goval/string"
	"github.com/remychantenay/goval/uuid"
	"reflect"
	"strings"
)

const (
	argTypeString      = "string"
	argTypeNumber      = "number"
	argTypeEmail       = "email"
	argTypeUUID        = "uuid"
	argTypeCountryCode = "country_code"
	argTypeCurrency    = "currency"
	argTypeEnum        = "enum"
)

// ValidateStruct validates any struct containing the nameTag goval
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

// getValidator returns the appropriate validator
func getValidator(tag string) generic.Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case argTypeNumber:
		return number.NewValidator(args[1:])
	case argTypeString:
		return str.NewValidator(args[1:])
	case argTypeEmail:
		return email.NewValidator(args[1:])
	case argTypeUUID:
		return uuid.NewValidator(args[1:])
	case argTypeCountryCode:
		return country.NewValidator(args[1:])
	case argTypeCurrency:
		return currency.NewValidator(args[1:])
	case argTypeEnum:
		return enum.NewValidator(args[1:])
	}

	return generic.GenericValidator{}
}
