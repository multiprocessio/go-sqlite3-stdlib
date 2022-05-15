package stdlib

import "testing"

func Test_regexp(t *testing.T) {
	assertQuery(t, "SELECT 'ZaB' REGEXP '[a-zA-Z]+'", "1")
	assertQuery(t, "SELECT 'ZaB0' REGEXP '[a-zA-Z]+$'", "0")

	// bad regexp
	assertQuery(t, "SELECT 'ZaB0' REGEXP ']]]]]'", "0")
}
