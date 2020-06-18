package email

import (
	"fmt"
	"github.com/remychantenay/goval/generic"
	"regexp"
	"strings"
	"sync"
)

const (
	// argConstraintMax
	argConstraintMax = "max="

	// argConstraintMin
	argConstraintMin = "min="

	// argConstraintRequired
	argConstraintRequired = "required="

	// argConstraintDomain (e.g. @google.com)
	argConstraintDomain = "domain="
)

type emailValidator struct {
	Min      int
	Max      int
	Required bool
	Domain   string
}

// regEx will be lazy loaded (via sync.Once)
var regEx *regexp.Regexp
var regExOnce sync.Once

func (v *emailValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required {
		return false, fmt.Errorf("cannot be nil")
	}

	str := val.(string)
	l := len(str)

	if l == 0 {
		if v.Required {
			return false, fmt.Errorf("cannot be blank")
		}
	} else {
		if v.Min != -1 && l < v.Min {
			return false, fmt.Errorf("should be at least %v characters long", v.Min)
		}
		if v.Max != -1 && v.Max >= v.Min && l > v.Max {
			return false, fmt.Errorf("should be less than %v characters long", v.Max)
		}

		// Lazy loading the regular expression
		regExOnceBody := func() {
			regEx = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
		}
		regExOnce.Do(regExOnceBody)

		if !regEx.MatchString(str) {
			return false, fmt.Errorf("is an invalid email address")
		}

		if len(v.Domain) != 0 && !strings.Contains(str, v.Domain) {
			return false, fmt.Errorf("is an invalid domain")
		}
	}

	return true, nil
}

// NewValidator builds the validator for email addresses
func NewValidator(args []string) generic.Validator {
	validator := emailValidator{-1, -1, false, ""}
	count := len(args) - 1
	for i := 0; i <= count; i++ {
		if strings.Contains(args[i], argConstraintMax) {
			fmt.Sscanf(args[i], argConstraintMax+"%d", &validator.Max)
		} else if strings.Contains(args[i], argConstraintMin) {
			fmt.Sscanf(args[i], argConstraintMin+"%d", &validator.Min)
		} else if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &validator.Required)
		} else if strings.Contains(args[i], argConstraintDomain) {
			fmt.Sscanf(args[i], argConstraintDomain+"%s", &validator.Domain)
		}
	}
	return &validator
}
