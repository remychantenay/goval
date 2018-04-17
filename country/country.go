package country

import (
	"fmt"
	"strings"
	"regexp"
	"github.com/remychantenay/goval/generic"
	"github.com/remychantenay/goval/constant"
)


type CountryCodeValidator struct {
	Required bool
	ExcludeEu bool
	Exclude string
}

func (v CountryCodeValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.Required { return false, fmt.Errorf("cannot be nil") }

	str := val.(string)
	l := len(str)
	expectedSize := 2

	if l == 0 {
		if v.Required {return false, fmt.Errorf("cannot be blank")}
	} else {

		if l != expectedSize {return false, fmt.Errorf("should be %v characters long", expectedSize)}

		if v.ExcludeEu {
			countryCodeRegEx := regexp.MustCompile("^(AF|AX|AL|DZ|AS|AD|AO|AI|AQ|AG|AR|AM|AW|AU|AT|AZ|BS|BH|BD|BB|BY|BE|BZ|BJ|BM|BT|BO|BQ|BA|BW|BV|BR|IO|BN|BG|BF|BI|KH|CM|CA|CV|KY|CF|TD|CL|CN|CX|CC|CO|KM|CG|CD|CK|CR|CI|HR|CU|CW|CY|CZ|DK|DJ|DM|DO|EC|EG|SV|GQ|ER|EE|ET|FK|FO|FJ|FI|FR|GF|PF|TF|GA|GM|GE|DE|GH|GI|GR|GL|GD|GP|GU|GT|GG|GN|GW|GY|HT|HM|VA|HN|HK|HU|IS|IN|ID|IR|IQ|IE|IM|IL|IT|JM|JP|JE|JO|KZ|KE|KI|KP|KR|KW|KG|LA|LV|LB|LS|LR|LY|LI|LT|LU|MO|MK|MG|MW|MY|MV|ML|MT|MH|MQ|MR|MU|YT|MX|FM|MD|MC|MN|ME|MS|MA|MZ|MM|NA|NR|NP|NL|NC|NZ|NI|NE|NG|NU|NF|MP|NO|OM|PK|PW|PS|PA|PG|PY|PE|PH|PN|PL|PT|PR|QA|RE|RO|RU|RW|BL|SH|KN|LC|MF|PM|VC|WS|SM|ST|SA|SN|RS|SC|SL|SG|SX|SK|SI|SB|SO|ZA|GS|SS|ES|LK|SD|SR|SJ|SZ|SE|CH|SY|TW|TJ|TZ|TH|TL|TG|TK|TO|TT|TN|TR|TM|TC|TV|UG|UA|AE|GB|US|UM|UY|UZ|VU|VE|VN|VG|VI|WF|EH|YE|ZM|ZW)$")
			if !countryCodeRegEx.MatchString(str) {return false, fmt.Errorf("is an invalid country code")}
		} else {
			countryCodeRegExNonEu := regexp.MustCompile("^(AF|AX|AL|DZ|AS|AD|AO|AI|AQ|AG|AR|AM|AW|AU|AZ|BS|BH|BD|BB|BY|BZ|BJ|BM|BT|BO|BQ|BA|BW|BV|BR|IO|BN|BF|BI|KH|CM|CA|CV|KY|CF|TD|CL|CN|CX|CC|CO|KM|CG|CD|CK|CR|CI|CU|CW|DJ|DM|DO|EC|EG|SV|GQ|ER|ET|FK|FO|FJ|GF|PF|TF|GA|GM|GE|GH|GI|GL|GD|GP|GU|GT|GG|GN|GW|GY|HT|HM|VA|HN|HK|IS|IN|ID|IR|IQ|IM|IL|JM|JP|JE|JO|KZ|KE|KI|KP|KR|KW|KG|LA|LB|LS|LR|LY|LI|MO|MK|MG|MW|MY|MV|ML|MH|MQ|MR|MU|YT|MX|FM|MD|MC|MN|ME|MS|MA|MZ|MM|NA|NR|NP|NC|NZ|NI|NE|NG|NU|NF|MP|NO|OM|PK|PW|PS|PA|PG|PY|PE|PH|PN|PR|QA|RE|RU|RW|BL|SH|KN|LC|MF|PM|VC|WS|SM|ST|SA|SN|RS|SC|SL|SG|SX|SB|SO|ZA|GS|SS|LK|SD|SR|SJ|SZ|CH|SY|TW|TJ|TZ|TH|TL|TG|TK|TO|TT|TN|TR|TM|TC|TV|UG|UA|AE|US|UM|UY|UZ|VU|VE|VN|VG|VI|WF|EH|YE|ZM|ZW)$")
			if !countryCodeRegExNonEu.MatchString(str) {return false, fmt.Errorf("is an invalid country code")}
		}

		b, err := generic.ValueExcluded(str, v.Exclude)
		if !b { return b, err }
	}

	return true, nil
}

// BuildCountryCodeValidator allows to build the validator for country codes (e.g. US)
func BuildCountryCodeValidator(args []string) generic.Validator {
	validator := CountryCodeValidator{false, false, ""}
	count := len(args)
	for i := 0; i < count; i++ {
		fmt.Println(args[i])
		if strings.Contains(args[i], constant.ArgConstraintRequired) {
			fmt.Sscanf(args[i], constant.ArgConstraintRequired+"%t", &validator.Required)
		} else if strings.Contains(args[i], constant.ArgConstraintExcludeEu) {
			fmt.Sscanf(args[i], constant.ArgConstraintExcludeEu+"%t", &validator.ExcludeEu)
		} else if strings.Contains(args[i], constant.ArgConstraintExclude) {
			fmt.Sscanf(args[i], constant.ArgConstraintExclude+"%s", &validator.Exclude)
		}
	}
	return validator
}

// BuildCountryCodeValidator allows to build the validator for country codes (e.g. US), mainly used for unit tests
func BuildCountryCodeValidatorWithFullTag(tag string) generic.Validator {
	args := strings.Split(tag, ",")
	return BuildCountryCodeValidator(args[1:])
}
