package stdlib

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func stringy(a any) string {
	if a == nil {
		return "null"
	}

	switch t := a.(type) {
	case string:
		return t
	case []byte:
		if len(t) == 0 {
			return "null"
		}

		return string(t)
	}

	return fmt.Sprintf("%v", a)
}

func stringy2int(f func(a, b string) int) func(a, b any) int {
	return func(a, b any) int {
		return f(stringy(a), stringy(b))
	}
}

func stringy1string(f func(a string) string) func(a any) string {
	return func(a any) string {
		return f(stringy(a))
	}
}

func stringy2string(f func(a, b string) string) func(a, b any) string {
	return func(a, b any) string {
		return f(stringy(a), stringy(b))
	}
}

func stringy3string(f func(a, b, c string) string) func(a, b, c any) string {
	return func(a, b, c any) string {
		return f(stringy(a), stringy(b), stringy(c))
	}
}

func ext_repeat(s any, n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(stringy(s))
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

func ext_lpad(_s any, length int, _padWith any) string {
	s := stringy(_s)

	if len(s) > length {
		return s[:length]
	}

	padWith := []rune(stringy(_padWith))

	var sb strings.Builder
	for i := 0; i < length-len(s); i++ {
		sb.WriteRune(padWith[i%len(padWith)])
	}

	sb.WriteString(s)
	return sb.String()[:length]
}

func ext_rpad(_s any, length int, _padWith any) string {
	s := stringy(_s)

	if len(s) > length {
		return s[:length]
	}

	padWith := []rune(stringy(_padWith))

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
	"strpos":    stringy2int(ext_charindex),
	"charindex": stringy2int(ext_charindex),
	"ltrim":     stringy2string(ext_ltrim),
	"rtrim":     stringy2string(ext_rtrim),
	"trim":      stringy2string(ext_ltrim),
	"replace":   stringy3string(strings.ReplaceAll),
	"reverse":   stringy1string(ext_reverse),
	"lpad":      ext_lpad,
	"rpad":      ext_rpad,
}
