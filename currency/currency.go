package currency

import (
	"fmt"
	"strings"
	"regexp"
	"github.com/remychantenay/goval/generic"
	"github.com/remychantenay/goval/constant"
)


type CurrencyValidator struct {
	Required bool
	Exclude string
}

func (v CurrencyValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	str := val.(string)
	l := len(str)
	expectedSize := 3

	if l == 0 {
		if v.Required {return false, fmt.Errorf("cannot be blank")}
	} else {
		if l != expectedSize {return false, fmt.Errorf("should be %v characters long", expectedSize)}

		currencyRegEx := regexp.MustCompile("/^AED|AFN|ALL|AMD|ANG|AOA|ARS|AUD|AWG|AZN|BAM|BBD|BDT|BGN|BHD|BIF|BMD|BND|BOB|BRL|BSD|BTN|BWP|BYR|BZD|CAD|CDF|CHF|CLP|CNY|COP|CRC|CUC|CUP|CVE|CZK|DJF|DKK|DOP|DZD|EGP|ERN|ETB|EUR|FJD|FKP|GBP|GEL|GGP|GHS|GIP|GMD|GNF|GTQ|GYD|HKD|HNL|HRK|HTG|HUF|IDR|ILS|IMP|INR|IQD|IRR|ISK|JEP|JMD|JOD|JPY|KES|KGS|KHR|KMF|KPW|KRW|KWD|KYD|KZT|LAK|LBP|LKR|LRD|LSL|LYD|MAD|MDL|MGA|MKD|MMK|MNT|MOP|MRO|MUR|MVR|MWK|MXN|MYR|MZN|NAD|NGN|NIO|NOK|NPR|NZD|OMR|PAB|PEN|PGK|PHP|PKR|PLN|PYG|QAR|RON|RSD|RUB|RWF|SAR|SBD|SCR|SDG|SEK|SGD|SHP|SLL|SOS|SPL|SRD|STD|SVC|SYP|SZL|THB|TJS|TMT|TND|TOP|TRY|TTD|TVD|TWD|TZS|UAH|UGX|USD|UYU|UZS|VEF|VND|VUV|WST|XAF|XCD|XDR|XOF|XPF|YER|ZAR|ZMW|ZWD$/")
		if !currencyRegEx.MatchString(str) {return false, fmt.Errorf("is an invalid currency")}

		b, err := generic.ValueExcluded(str, v.Exclude)
		if !b {return b, err }
	}

	return true, nil
}

// BuildCurrencyValidator allows to build the validator for currency codes (e.g. EUR)
func BuildCurrencyValidator(args []string) generic.Validator {
	validator := CurrencyValidator{false, ""}
	count := len(args)-1
	for i := 0; i <= count; i++ {
		if strings.Contains(args[i], constant.ArgConstraintRequired) {
			fmt.Sscanf(args[i], constant.ArgConstraintRequired+"%t", &validator.Required)
		} else if strings.Contains(args[i], constant.ArgConstraintExclude) {
			fmt.Sscanf(args[i], constant.ArgConstraintExclude+"%s", &validator.Exclude)
		}
	}
	return validator
}

// BuildCurrencyValidatorWithFullTag allows to build the validator for currency codes (e.g. EUR), mainly used for unit tests
func BuildCurrencyValidatorWithFullTag(tag string) generic.Validator {
	args := strings.Split(tag, ",")
	return BuildCurrencyValidator(args[1:])
}
