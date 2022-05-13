package stdlib

import "testing"

func Test_ext_regexp(t *testing.T) {
	assertQuery(t, "SELECT 'ZaB' REGEXP '[a-zA-Z]+'", "1")
	assertQuery(t, "SELECT 'ZaB0' REGEXP '[a-zA-Z]+$'", "0")
}
