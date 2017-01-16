package generated

import (
	"fmt"
	"regexp"
	"strings"
)

var FormatValidators = map[string]func(interface{}) (bool, string){
	"uuid": func(value interface{}) (bool, string) {
		stringValue, ok := value.(string)
		if !ok {
			return false, fmt.Sprintf("%v is not a valid uuid", value)
		}
		stringValue = strings.ToLower(stringValue)
		r := regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|aA|bB][a-f0-9]{3}-[a-f0-9]{12}$")
		if !r.MatchString(stringValue) {
			return false, fmt.Sprintf("%v is not a valid uuid", value)
		}

		return true, ""
	},
}

