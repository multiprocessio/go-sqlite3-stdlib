package stdlib

import (
	"regexp"
)

func _regexp(re, s string) bool {
	c, err := regexp.Compile(re)
	if err != nil {
		return false
	}

	return c.MatchString(s)
}

var regexpFunctions = map[string]any{
	"regexp": _regexp,
}
