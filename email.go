package goval

import (
	"fmt"
	"strings"
	"regexp"
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

	if l == 0  && v.Required {return false, fmt.Errorf("cannot be blank")}
	if v.Min != -1 && l < v.Min {return false, fmt.Errorf("should be at least %v characters long", v.Min)}
	if v.Max != -1 && v.Max >= v.Min && l > v.Max {return false, fmt.Errorf("should be less than %v characters long", v.Max)}

	var mailRegEx = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
	if !mailRegEx.MatchString(str) {return false, fmt.Errorf("invalid email address")}

	if strings.Compare(v.Domain, "") != 0 && !strings.Contains(str, v.Domain) {return false, fmt.Errorf("invalid domain")}

	return true, nil
}

func buildEmailValidator(args []string) Validator {
	validator := EmailValidator{-1,-1,false, ""}
	count := len(args)
	for i := 0; i <= count; i++ {
		fmt.Println(args[i])
		if strings.Contains(args[i], ArgConstraintMax) {
			fmt.Sscanf(args[i], ArgConstraintMax+"%d", &validator.Max)
		} else if strings.Contains(args[i], ArgConstraintMin) {
			fmt.Sscanf(args[i], ArgConstraintMin+"%d", &validator.Min)
		} else if strings.Contains(args[i], ArgConstraintRequired) {
			fmt.Sscanf(args[i], ArgConstraintRequired+"%t", &validator.Required)
		} else if strings.Contains(args[i], ArgConstraintDomain) {
			fmt.Sscanf(args[i], ArgConstraintDomain+"%t", &validator.Domain)
		}
	}
	return validator
}
