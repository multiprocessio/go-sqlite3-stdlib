package runner

import (
	"regexp"
)

func ext_regexp(re, s string) bool {
	matched, err := regexp.MatchString(re, s)
	if err != nil {
		return false
	}

	return matched
}
