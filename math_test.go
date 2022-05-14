package stdlib

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

var initialized = false

func assertQueryPrepare(t *testing.T, queries []string, res string) {
	if !initialized {
		Register("sqlite3_ext")
		initialized = true
	}

	db, err := sql.Open("sqlite3_ext", ":memory:")
	assert.Nil(t, err)

	for _, query := range queries[:len(queries)-1] {
		_, err = db.Exec(query)
		assert.Nil(t, err)
	}

	query := queries[len(queries)-1]
	t.Logf("Query: %s, expecting: %s", query, res)
	var s string
	err = db.QueryRow(query).Scan(&s)
	assert.Nil(t, err)
	assert.Equal(t, res, s)
}

func assertQuery(t *testing.T, query string, res string) {
	assertQueryPrepare(t, []string{query}, res)
}

func Test_floaty(t *testing.T) {
	tests := []struct {
		in  any
		out float64
	}{
		{"0.1", 0.1},
		{"sdflkj", 0},
		{12, 12.0},
		{nil, 0},
		{int8(1), 1},
		{int16(1), 1},
		{int32(1), 1},
		{int64(1), 1},
		{uint(1), 1},
		{uint16(1), 1},
		{uint32(1), 1},
		{uint64(1), 1},
		{float32(1), 1},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, floaty(test.in))
	}
}

func Test_acos(t *testing.T) {
	assertQuery(t, "SELECT acos(0.1)", "1.4706289056333368")
}

func Test_radians(t *testing.T) {
	assertQuery(t, "SELECT radians(180)", "3.141592653589793")
}

func Test_degrees(t *testing.T) {
	assertQuery(t, "SELECT degrees(3.141592653589793)", "180")
}

func Test_mod(t *testing.T) {
	assertQuery(t, "SELECT mod('10', '2')", "0")
}

func Test_pi(t *testing.T) {
	assertQuery(t, "SELECT pi()", "3.141592653589793")
}

func Test_acosh(t *testing.T) {
	assertQuery(t, "SELECT acosh(2.0)", "1.3169578969248166")
}

func Test_trunc(t *testing.T) {
	assertQuery(t, "SELECT trunc(10.4)", "10")
	assertQuery(t, "SELECT trunc(-10.9)", "-10")
	assertQuery(t, "SELECT trunc(0.49999)", "0")
	assertQuery(t, "SELECT trunc(-0.49999)", "0")
}

func Test_floor(t *testing.T) {
	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n INT)",
		"INSERT INTO x VALUES (2)",
		"SELECT floor(n) FROM x",
	}, "2")
}
