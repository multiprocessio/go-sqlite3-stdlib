package stdlib

import (
	"strings"
	"unicode/utf8"
)

func ext_repeat(s string, n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(s)
	}

	return sb.String()
}

func ext_charindex(s, sub string) int {
	return strings.Index(s, sub)
}

func ext_ltrim(s, characters string) string {
	for {
		old := s

		for _, c := range characters {
			s = strings.TrimLeft(s, string(c))
		}

		// Keep on trimming while any character appears on the left.
		// e.g. ltrim('abcabcd', 'abc') == 'd'
		if s == old {
			break
		}
	}

	return s
}

func ext_rtrim(s, characters string) string {
	for {
		old := s

		for _, c := range characters {
			s = strings.TrimRight(s, string(c))
		}

		// Keep on trimming while any character appears on the left.
		// e.g. ltrim('abcabcd', 'abc') == 'd'
		if s == old {
			break
		}
	}

	return s
}

// SOURCE: https://stackoverflow.com/a/20225618/1507139
func ext_reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func ext_lpad(s string, length int, padWidthS string) string {
	if len(s) > length {
		return s[:length]
	}

	padWith := []rune(padWidthS)

	var sb strings.Builder
	for i := 0; i < length-len(s); i++ {
		sb.WriteRune(padWith[i%len(padWith)])
	}

	sb.WriteString(s)
	return sb.String()[:length]
}

func ext_rpad(s string, length int, options ...any) string {
	if len(s) > length {
		return s[:length]
	}

	padWith := []rune(" ")
	if len(options) > 0 {
		if p, ok := options[0].(string); ok {
			padWith = []rune(p)
		}
	}

	var sb strings.Builder
	sb.WriteString(s)
	for i := 0; i < length-len(s); i++ {
		sb.WriteRune(padWith[i%len(padWith)])
	}

	return sb.String()
}

var stringFunctions = map[string]any{
	"repeat":    ext_repeat,
	"replicate": ext_repeat,
	"strpos":    ext_charindex,
	"charindex": ext_charindex,
	"ltrim":     ext_ltrim,
	"rtrim":     ext_rtrim,
	"trim":      ext_ltrim,
	"replace":   strings.ReplaceAll,
	"reverse":   ext_reverse,
	"lpad":      ext_lpad,
	"rpad":      ext_rpad,
	"upper":     strings.ToUpper,
	"lower":     strings.ToLower,
}
