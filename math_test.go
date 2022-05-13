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

func Test_ext_acos(t *testing.T) {
	assertQuery(t, "SELECT acos(0.1)", "1.4706289056333368")
}

func Test_ext_acosh(t *testing.T) {
	assertQuery(t, "SELECT acosh(2.0)", "1.3169578969248166")
}

func Test_ext_trunc(t *testing.T) {
	assertQuery(t, "SELECT trunc(10.4)", "10")
	assertQuery(t, "SELECT trunc(-10.9)", "-10")
	assertQuery(t, "SELECT trunc(0.49999)", "0")
	assertQuery(t, "SELECT trunc(-0.49999)", "0")
}
