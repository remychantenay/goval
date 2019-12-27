package email

import (
	"fmt"
	"strings"
	"regexp"
	"github.com/remychantenay/goval/generic"
	"github.com/remychantenay/goval/constant"
)


type EmailValidator struct {
	Min int
	Max int
	Required bool
	Domain string
}

func (v EmailValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	str := val.(string)
	l := len(str)

	if l == 0 {
		if v.Required {return false, fmt.Errorf("cannot be blank")}
	} else {
		if v.Min != -1 && l < v.Min {return false, fmt.Errorf("should be at least %v characters long", v.Min)}
		if v.Max != -1 && v.Max >= v.Min && l > v.Max {return false, fmt.Errorf("should be less than %v characters long", v.Max)}

		var mailRegEx = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
		if !mailRegEx.MatchString(str) {return false, fmt.Errorf("is an invalid email address")}

		if len(v.Domain) != 0 && !strings.Contains(str, v.Domain) {return false, fmt.Errorf("is an invalid domain")}
	}

	return true, nil
}

// BuildEmailValidator allows to build the validator for email addresses
func BuildEmailValidator(args []string) generic.Validator {
	validator := EmailValidator{-1,-1,false, ""}
	count := len(args)-1
	for i := 0; i <= count; i++ {
		if strings.Contains(args[i], constant.ArgConstraintMax) {
			fmt.Sscanf(args[i], constant.ArgConstraintMax+"%d", &validator.Max)
		} else if strings.Contains(args[i], constant.ArgConstraintMin) {
			fmt.Sscanf(args[i], constant.ArgConstraintMin+"%d", &validator.Min)
		} else if strings.Contains(args[i], constant.ArgConstraintRequired) {
			fmt.Sscanf(args[i], constant.ArgConstraintRequired+"%t", &validator.Required)
		} else if strings.Contains(args[i], constant.ArgConstraintDomain) {
			fmt.Sscanf(args[i],constant.ArgConstraintDomain+"%s", &validator.Domain)
		}
	}
	return validator
}

// BuildEmailValidator allows to build the validator for email addresses, mainly used for unit tests
func BuildEmailValidatorWithFullTag(tag string) generic.Validator {
	args := strings.Split(tag, ",")
	return BuildEmailValidator(args[1:])
}