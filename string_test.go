package stdlib

import (
	"fmt"
	"testing"
)

func Test_ext_repeat(t *testing.T) {
	assertQuery(t, "SELECT repeat('a', 3)", "aaa")
	assertQuery(t, "SELECT replicate('a', 3)", "aaa")
}

func Test_ext_strpos(t *testing.T) {
	tests := []struct {
		haystack      string
		needle        string
		expectedIndex string
	}{
		{
			"abcde",
			"f",
			"-1",
		},
		{
			"abcde",
			"e",
			"4",
		},
		{
			"abcde",
			"c",
			"2",
		},
		{
			"abcde",
			"a",
			"0",
		},
	}

	for _, test := range tests {
		assertQuery(t, fmt.Sprintf("SELECT strpos('%s', '%s')", test.haystack, test.needle), test.expectedIndex)
		assertQuery(t, fmt.Sprintf("SELECT charindex('%s', '%s')", test.haystack, test.needle), test.expectedIndex)
	}
}

func Test_ext_ltrim(t *testing.T) {
	tests := []struct {
		in   string
		trim string
		out  string
	}{
		{
			"abcde",
			"abc",
			"de",
		},
		{
			"abcbcade",
			"abc",
			"de",
		},
		{
			"abc",
			"",
			"abc",
		},
		{
			"abcx",
			"x",
			"abcx",
		},
	}

	for _, test := range tests {
		assertQuery(t, fmt.Sprintf("SELECT ltrim('%s', '%s');", test.in, test.trim), test.out)
		assertQuery(t, fmt.Sprintf("SELECT trim('%s', '%s');", test.in, test.trim), test.out)
	}
}

func Test_ext_rtrim(t *testing.T) {
	tests := []struct {
		in   string
		trim string
		out  string
	}{
		{
			"abceedde",
			"de",
			"abc",
		},
		{
			"abcbcade",
			"abc",
			"abcbcade",
		},
	}

	for _, test := range tests {
		assertQuery(t, fmt.Sprintf("SELECT rtrim('%s', '%s');", test.in, test.trim), test.out)
	}
}

func Test_ext_replace(t *testing.T) {
	assertQuery(t, "SELECT replace('  whatever ', 'whatever', 'blah')", "  blah ")
}

func Test_ext_reverse(t *testing.T) {
	assertQuery(t, "SELECT reverse('bubbly')", "ylbbub")
}

func Test_ext_lpad(t *testing.T) {
	tests := []struct {
		in     string
		length int
		pad    string
		out    string
	}{
		{
			"abcde",
			3,
			"0",
			"abc",
		},
		{
			"aa",
			3,
			"0",
			"0aa",
		},
		{
			"aa",
			2,
			"0",
			"aa",
		},
		{
			"a",
			0,
			"0",
			"",
		},
	}

	for _, test := range tests {
		assertQuery(t, fmt.Sprintf("SELECT lpad('%s', %d, '%s');", test.in, test.length, test.pad), test.out)
	}

	// Test no third argument variation
	assertQuery(t, "SELECT lpad('22', 3, '0');", "022")
}

func Test_ext_rpad(t *testing.T) {
	tests := []struct {
		in     string
		length int
		pad    string
		out    string
	}{
		{
			"abcde",
			3,
			"0",
			"abc",
		},
		{
			"aa",
			3,
			"0",
			"aa0",
		},
		{
			"aa",
			2,
			"0",
			"aa",
		},
		{
			"a",
			0,
			"0",
			"",
		},
	}

	for _, test := range tests {
		assertQuery(t, fmt.Sprintf("SELECT rpad('%s', %d, '%s');", test.in, test.length, test.pad), test.out)
	}

	// Test no 3rd argument variant
	assertQuery(t, "SELECT rpad('22', 3, '0');", "220")
}
