package email

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
	argConstraintRequired = "required="

	// argConstraintMax defines the maximum number of characters allowed.
	argConstraintMax = "max="

	// argConstraintMin defines the minimum number of characters allowed.
	argConstraintMin = "min="

	// argConstraintDomain defines if only one domain is expected
	// (e.g. @google.com)
	argConstraintDomain = "domain="
)

var pool = sync.Pool{
	New: func() interface{} { return &emailValidator{-1, -1, false, ""} },
}

type emailValidator struct {
	min      int
	max      int
	required bool
	domain   string
}

// regEx will be lazy loaded (via sync.Once)
var regEx *regexp.Regexp
var regExOnce sync.Once

func (v *emailValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.required {
		return false, fmt.Errorf("cannot be nil")
	}

	str := val.(string)
	l := len(str)

	if l == 0 {
		if v.required {
			return false, fmt.Errorf("cannot be blank")
		}
	} else {
		if v.min != -1 && l < v.min {
			return false, fmt.Errorf("should be at least %d characters long", v.min)
		}
		if v.max != -1 && v.max >= v.min && l > v.max {
			return false, fmt.Errorf("should be less than %d characters long", v.max)
		}

		// Lazy loading the regular expression
		regExOnceBody := func() {
			regEx = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
		}
		regExOnce.Do(regExOnceBody)

		if !regEx.MatchString(str) {
			return false, fmt.Errorf("is an invalid email address")
		}

		if len(v.domain) != 0 && !strings.Contains(str, v.domain) {
			return false, fmt.Errorf("is an invalid domain")
		}
	}

	return true, nil
}

// NewValidator builds the validator for email addresses
func NewValidator(args []string) generic.Validator {
	validator := pool.Get().(*emailValidator)
	defer pool.Put(validator)
	if validator.max != -1 {
		validator.min, validator.max = -1, -1
		validator.required = false
		validator.domain = ""
	}

	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], argConstraintMax) {
			fmt.Sscanf(args[i], argConstraintMax+"%d", &validator.max)
		} else if strings.Contains(args[i], argConstraintMin) {
			fmt.Sscanf(args[i], argConstraintMin+"%d", &validator.min)
		} else if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &validator.required)
		} else if strings.Contains(args[i], argConstraintDomain) {
			fmt.Sscanf(args[i], argConstraintDomain+"%s", &validator.domain)
		}
	}
	return validator
}
