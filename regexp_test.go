package stdlib

import "testing"

func Test_regexp(t *testing.T) {
	assertQuery(t, "SELECT 'ZaB' REGEXP '[a-zA-Z]+'", "1")
	assertQuery(t, "SELECT 'ZaB0' REGEXP '[a-zA-Z]+$'", "0")

	// bad regexp
	assertQuery(t, "SELECT 'ZaB0' REGEXP ']['", "0")
}

func Test_regexp_split_part(t *testing.T) {
	assertQuery(t, "SELECT regexp_split_part('ab12', '[a-zA-Z]1', 1)", "2")
	assertQuery(t, "SELECT regexp_split_part('ab12', '[a-zA-Z]1', -1)", "2")
	assertQuery(t, "SELECT regexp_split_part('ab12', '[a-zA-Z]1', 100)", "")

	// bad regexp
	assertQuery(t, "SELECT regexp_split_part('ab12', '][', 1)", "")
}

func Test_regexp_count(t *testing.T) {
	assertQuery(t, "SELECT regexp_count('ab12', '[a-zA-Z]1')", "1")
	assertQuery(t, "SELECT regexp_count('ac22', '[a-zA-Z]1')", "0")

	// bad regexp
	assertQuery(t, "SELECT regexp_count('ab12', '][')", "0")
}
