package country

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/remychantenay/goval/generic"
)

const (
	// argConstraintRequired
	// (i.e. true or false)
	argConstraintRequired = "required="

	// argConstraintexcludeEU
	// (i.e. true or false)
	argConstraintexcludeEU = "exclude_eu="

	// argConstraintExclude - must be separated w/ a pipe (e.g. exclude=EUR|GBP)
	argConstraintExclude = "exclude="
)

var pool = sync.Pool{
	New: func() interface{} { return &countryCodeValidator{false, false, ""} },
}

type countryCodeValidator struct {
	required  bool
	excludeEU bool
	exclude   string
}

// RegularExpressions will be lazy loaded (via sync.Once)
type RegularExpressions struct {
	all       *regexp.Regexp
	withoutEU *regexp.Regexp
}

var countryCodeRegularExpressions RegularExpressions
var loadRegularExpressionsOnce sync.Once

// Validate validates the given struct
//
// Returns true if valid, false otherwise
func (v *countryCodeValidator) Validate(val interface{}) (bool, error) {
	if val == nil && v.required {
		return false, fmt.Errorf("cannot be nil")
	}

	strVal := val.(string)
	expectedSize := 2

	if len(strVal) == 0 && v.required {
		return false, fmt.Errorf("cannot be blank")
	} else {
		if len(strVal) != expectedSize {
			return false, fmt.Errorf("should be %d characters long", expectedSize)
		}

		// Lazy loading the regular expressions
		loadRegularExpressionsOnceBody := func() {
			countryCodeRegularExpressions = RegularExpressions{
				all:       regexp.MustCompile("^(AF|AX|AL|DZ|AS|AD|AO|AI|AQ|AG|AR|AM|AW|AU|AT|AZ|BS|BH|BD|BB|BY|BE|BZ|BJ|BM|BT|BO|BQ|BA|BW|BV|BR|IO|BN|BG|BF|BI|KH|CM|CA|CV|KY|CF|TD|CL|CN|CX|CC|CO|KM|CG|CD|CK|CR|CI|HR|CU|CW|CY|CZ|DK|DJ|DM|DO|EC|EG|SV|GQ|ER|EE|ET|FK|FO|FJ|FI|FR|GF|PF|TF|GA|GM|GE|DE|GH|GI|GR|GL|GD|GP|GU|GT|GG|GN|GW|GY|HT|HM|VA|HN|HK|HU|IS|IN|ID|IR|IQ|IE|IM|IL|IT|JM|JP|JE|JO|KZ|KE|KI|KP|KR|KW|KG|LA|LV|LB|LS|LR|LY|LI|LT|LU|MO|MK|MG|MW|MY|MV|ML|MT|MH|MQ|MR|MU|YT|MX|FM|MD|MC|MN|ME|MS|MA|MZ|MM|NA|NR|NP|NL|NC|NZ|NI|NE|NG|NU|NF|MP|NO|OM|PK|PW|PS|PA|PG|PY|PE|PH|PN|PL|PT|PR|QA|RE|RO|RU|RW|BL|SH|KN|LC|MF|PM|VC|WS|SM|ST|SA|SN|RS|SC|SL|SG|SX|SK|SI|SB|SO|ZA|GS|SS|ES|LK|SD|SR|SJ|SZ|SE|CH|SY|TW|TJ|TZ|TH|TL|TG|TK|TO|TT|TN|TR|TM|TC|TV|UG|UA|AE|GB|US|UM|UY|UZ|VU|VE|VN|VG|VI|WF|EH|YE|ZM|ZW)$"),
				withoutEU: regexp.MustCompile("^(AF|AX|AL|DZ|AS|AD|AO|AI|AQ|AG|AR|AM|AW|AU|AZ|BS|BH|BD|BB|BY|BZ|BJ|BM|BT|BO|BQ|BA|BW|BV|BR|IO|BN|BF|BI|KH|CM|CA|CV|KY|CF|TD|CL|CN|CX|CC|CO|KM|CG|CD|CK|CR|CI|CU|CW|DJ|DM|DO|EC|EG|SV|GQ|ER|ET|FK|FO|FJ|GF|PF|TF|GA|GM|GE|GH|GI|GL|GD|GP|GU|GT|GG|GN|GW|GY|HT|HM|VA|HN|HK|IS|IN|ID|IR|IQ|IM|IL|JM|JP|JE|JO|KZ|KE|KI|KP|KR|KW|KG|LA|LB|LS|LR|LY|LI|MO|MK|MG|MW|MY|MV|ML|MH|MQ|MR|MU|YT|MX|FM|MD|MC|MN|ME|MS|MA|MZ|MM|NA|NR|NP|NC|NZ|NI|NE|NG|NU|NF|MP|NO|OM|PK|PW|PS|PA|PG|PY|PE|PH|PN|PR|QA|RE|RU|RW|BL|SH|KN|LC|MF|PM|VC|WS|SM|ST|SA|SN|RS|SC|SL|SG|SX|SB|SO|ZA|GS|SS|LK|SD|SR|SJ|SZ|CH|SY|TW|TJ|TZ|TH|TL|TG|TK|TO|TT|TN|TR|TM|TC|TV|UG|UA|AE|US|UM|UY|UZ|VU|VE|VN|VG|VI|WF|EH|YE|ZM|ZW)$"),
			}
		}
		loadRegularExpressionsOnce.Do(loadRegularExpressionsOnceBody)

		if v.excludeEU {
			if !countryCodeRegularExpressions.all.MatchString(strVal) {
				return false, fmt.Errorf("is an invalid country code")
			}
		} else {
			if !countryCodeRegularExpressions.withoutEU.MatchString(strVal) {
				return false, fmt.Errorf("is an invalid country code")
			}
		}

		b, err := generic.ValueExcluded(strVal, v.exclude)
		if !b {
			return b, err
		}
	}

	return true, nil
}

// NewValidator build and return the validator for country codes (e.g. US)
func NewValidator(args []string) generic.Validator {
	v := pool.Get().(*countryCodeValidator)
	defer pool.Put(v)
	if !v.required || v.exclude != "" || !v.excludeEU {
		v.required = false
		v.exclude = ""
		v.excludeEU = false
	}
	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], argConstraintRequired) {
			fmt.Sscanf(args[i], argConstraintRequired+"%t", &v.required)
		} else if strings.Contains(args[i], argConstraintexcludeEU) {
			fmt.Sscanf(args[i], argConstraintexcludeEU+"%t", &v.excludeEU)
		} else if strings.Contains(args[i], argConstraintExclude) {
			fmt.Sscanf(args[i], argConstraintExclude+"%s", &v.exclude)
		}
	}
	return v
}
