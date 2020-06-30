package currency

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/remychantenay/goval/generic"
)

const (

	// argConstraintRequired defines whether or not the field is required.
	// In this context, required means not nil and not blank.
	// (true or false)
	argConstraintRequired = "required="

	// argConstraintExclude - must separated w/ a pipe
	// (e.g. exclude=EUR|GBP)
	argConstraintExclude = "exclude="

	// expectedLength - ISO-4217
	expectedLength = 3
)

var pool = sync.Pool{
	New: func() interface{} { return &currencyValidator{false, ""} },
}

type currencyValidator struct {
	required bool
	exclude  string
}

// regEx will be lazy loaded (via sync.Once)
var regEx *regexp.Regexp
var regExOnce sync.Once

// Validate the given value
// Must comply with ISO-4217 (see https://en.wikipedia.org/wiki/ISO_4217)
func (v *currencyValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.required {
		return false, fmt.Errorf("cannot be nil")
	}

	str := val.(string)

	if len(str) == 0 && v.required {
		return false, fmt.Errorf("cannot be blank")
	}
	if len(str) != expectedLength {
		return false, fmt.Errorf("should be %d characters long", expectedLength)
	}

	// Lazy loading the regular expression
	regExOnceBody := func() {
		regEx = regexp.MustCompile("/^AED|AFN|ALL|AMD|ANG|AOA|ARS|AUD|AWG|AZN|BAM|BBD|BDT|BGN|BHD|BIF|BMD|BND|BOB|BRL|BSD|BTN|BWP|BYR|BZD|CAD|CDF|CHF|CLP|CNY|COP|CRC|CUC|CUP|CVE|CZK|DJF|DKK|DOP|DZD|EGP|ERN|ETB|EUR|FJD|FKP|GBP|GEL|GGP|GHS|GIP|GMD|GNF|GTQ|GYD|HKD|HNL|HRK|HTG|HUF|IDR|ILS|IMP|INR|IQD|IRR|ISK|JEP|JMD|JOD|JPY|KES|KGS|KHR|KMF|KPW|KRW|KWD|KYD|KZT|LAK|LBP|LKR|LRD|LSL|LYD|MAD|MDL|MGA|MKD|MMK|MNT|MOP|MRO|MUR|MVR|MWK|MXN|MYR|MZN|NAD|NGN|NIO|NOK|NPR|NZD|OMR|PAB|PEN|PGK|PHP|PKR|PLN|PYG|QAR|RON|RSD|RUB|RWF|SAR|SBD|SCR|SDG|SEK|SGD|SHP|SLL|SOS|SPL|SRD|STD|SVC|SYP|SZL|THB|TJS|TMT|TND|TOP|TRY|TTD|TVD|TWD|TZS|UAH|UGX|USD|UYU|UZS|VEF|VND|VUV|WST|XAF|XCD|XDR|XOF|XPF|YER|ZAR|ZMW|ZWD$/")
	}
	regExOnce.Do(regExOnceBody)

	if !regEx.MatchString(str) {
		return false, fmt.Errorf("is an invalid currency")
	}

	b, err := generic.ValueExcluded(str, v.exclude)
	if !b {
		return b, err
	}

	return true, nil
}

// NewValidator build and return the validator for currency codes (e.g. EUR)
func NewValidator(args []string) generic.Validator {
	validator := pool.Get().(*currencyValidator)
	defer pool.Put(validator)
	if !validator.required || validator.exclude != "" {
		validator.required = false
		validator.exclude = ""
	}

	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &validator.required)
		} else if strings.Contains(args[i], argConstraintExclude) {
			fmt.Sscanf(args[i], argConstraintExclude+"%s", &validator.exclude)
		}
	}
	return validator
}
