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

func stringy1int64(f func(a string) int64) func(a any) int64 {
	return func(a any) int64 {
		return f(stringy(a))
	}
}

func stringy2int64(f func(a, b string) int64) func(a, b any) int64 {
	return func(a, b any) int64 {
		return f(stringy(a), stringy(b))
	}
}

func stringy1string(f func(a string) string) func(a any) string {
	return func(a any) string {
		return f(stringy(a))
	}
}

func repeat(s any, n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(stringy(s))
	}

	return sb.String()
}

func charindex(s, sub string) int64 {
	return int64(strings.Index(s, sub))
}

// SOURCE: https://stackoverflow.com/a/20225618/1507139
func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func lpad(_s any, length int, _padWith ...any) string {
	s := stringy(_s)

	if len(s) > length {
		return s[:length]
	}

	padWith := []rune(" ")
	if len(_padWith) > 0 {
		padWith = []rune(stringy(_padWith[0]))
	}

	var sb strings.Builder
	for i := 0; i < length-len(s); i++ {
		sb.WriteRune(padWith[i%len(padWith)])
	}

	sb.WriteString(s)
	return sb.String()[:length]
}

func rpad(_s any, length int, _padWith ...any) string {
	s := stringy(_s)

	if len(s) > length {
		return s[:length]
	}

	padWith := []rune(" ")
	if len(_padWith) > 0 {
		padWith = []rune(stringy(_padWith[0]))
	}

	var sb strings.Builder
	sb.WriteString(s)
	for i := 0; i < length-len(s); i++ {
		sb.WriteRune(padWith[i%len(padWith)])
	}

	return sb.String()
}

func length(s string) int64 {
	return int64(len(s))
}

func splitPart(s, sub, _part any) string {
	part := int(floaty(_part))
	parts := strings.Split(stringy(s), stringy(sub))
	if len(parts) == 0 || part >= len(parts) {
		return ""
	}

	if part < 0 {
		part = (len(parts) + part) % len(parts)
	}

	return parts[part]
}

var stringFunctions = map[string]any{
	"len":        stringy1int64(length),
	"split_part": splitPart,
	"repeat":     repeat,
	"replicate":  repeat,
	"strpos":     stringy2int64(charindex),
	"charindex":  stringy2int64(charindex),
	"reverse":    stringy1string(reverse),
	"lpad":       lpad,
	"rpad":       rpad,
}
