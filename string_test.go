package stdlib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_string(t *testing.T) {
	assert.Equal(t, "abc", stringy([]byte("abc")))
	assert.Equal(t, "null", stringy(nil))
}

func Test_repeat(t *testing.T) {
	assertQuery(t, "SELECT repeat('a', 3)", "aaa")
	assertQuery(t, "SELECT replicate('a', 3)", "aaa")
	assertQuery(t, "SELECT repeat(3, 3)", "333")
	assertQuery(t, "SELECT repeat(null, 3)", "nullnullnull")
}

func Test_strpos(t *testing.T) {
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

	assertQuery(t, "SELECT strpos(1234, 3)", "2")
}

func Test_ltrim(t *testing.T) {
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
	}

	assertQuery(t, "SELECT ltrim(' ac');", "ac")
}

func Test_rtrim(t *testing.T) {
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

	assertQuery(t, "SELECT rtrim(3334, 4);", "333")
	assertQuery(t, "SELECT rtrim(' ac ');", " ac")
}

func Test_trim(t *testing.T) {
	assertQuery(t, "SELECT trim('  whatever ')", "whatever")
	assertQuery(t, "SELECT trim('xabcx', 'x')", "abc")
}

func Test_replace(t *testing.T) {
	assertQuery(t, "SELECT replace('  whatever ', 'whatever', 'blah')", "  blah ")
	assertQuery(t, "SELECT replace(3443, 44, 55)", "3553")
}

func Test_reverse(t *testing.T) {
	assertQuery(t, "SELECT reverse('234')", "432")
}

func Test_lpad(t *testing.T) {
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
	assertQuery(t, "SELECT lpad('22', 3);", " 22")

	// int variation
	assertQuery(t, "SELECT lpad(22, 3, 0);", "022")
}

func Test_rpad(t *testing.T) {
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
	assertQuery(t, "SELECT rpad('22', 3);", "22 ")
}
