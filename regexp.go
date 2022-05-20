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

func regexpSplitPart(_s, _split, _part any) string {
	s := stringy(_s)
	split := stringy(_split)
	part := int(floaty(_part))

	c, err := regexp.Compile(split)
	if err != nil {
		return ""
	}

	pieces := part + 1
	if pieces == 0 {
		pieces = -1
	}

	parts := c.Split(s, pieces)

	if len(parts) == 0 || part >= len(parts) {
		return ""
	}

	if part < 0 {
		part = (len(parts) + part) % len(parts)
	}

	return parts[part]
}

func regexpCount(_s, _re any) int64 {
	s := stringy(_s)
	re := stringy(_re)

	c, err := regexp.Compile(re)
	if err != nil {
		return 0
	}

	return int64(len(c.FindAllStringIndex(s, -1)))
}

var regexpFunctions = map[string]any{
	"regexp":            _regexp,
	"regexp_split_part": regexpSplitPart,
	"regexp_count":      regexpCount,
}
