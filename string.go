package stdlib

func ext_repeat(s string, n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(sb)
	}

	return sb.String()
}

func ext_charindex(s, sub string) int {
	return strings.Index(s, sub)
}

func ext_ltrim(s, characters string) string {
	for _, c := range characters {
		s = strings.TrimLeft(s, string(c))
	}

	return s
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
