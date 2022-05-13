package stdlib

import "testing"

func Test_stddev(t *testing.T) {
	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n INT)",
		"INSERT INTO x VALUES (1), (2)",
		"SELECT stddev_pop(n) FROM x",
	}, "0.5")
}

func Test_mode(t *testing.T) {
	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n INT)",
		"INSERT INTO x VALUES (1), (2), (2)",
		"SELECT mode(n) FROM x",
	}, "2")

	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n TEXT)",
		"INSERT INTO x VALUES ('a'), ('b'), ('a')",
		"SELECT modestr(n) FROM x",
	}, "a")
}
