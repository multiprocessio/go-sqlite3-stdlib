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
		"SELECT mode(n) FROM x",
	}, "a")
}

func Test_median(t *testing.T) {
	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n INT)",
		"INSERT INTO x VALUES (1), (2), (2)",
		"SELECT median(n) FROM x",
	}, "2")

	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n INT)",
		"INSERT INTO x VALUES (1), (2), (3)",
		"SELECT median(n) FROM x",
	}, "2")

	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n INT)",
		"INSERT INTO x VALUES (1)",
		"SELECT median(n) FROM x",
	}, "1")

	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n INT)",
		"SELECT coalesce(median(n), 'null') FROM x",
	}, "null")

	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n TEXT)",
		"INSERT INTO x VALUES ('a'), ('b'), ('a'), ('c'), ('d')",
		"SELECT median(n) FROM x",
	}, "b")

	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n TEXT)",
		"INSERT INTO x VALUES (null), (null)",
		"SELECT coalesce(median(n), 'null') FROM x",
	}, "null")

	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n REAL)",
		"INSERT INTO x VALUES (1.2), (3.4), (4.4)",
		"SELECT median(n) FROM x",
	}, "3.4")
}

func Test_percentile(t *testing.T) {
	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n INT)",
		"INSERT INTO x VALUES (1), (3), (4)",
		"SELECT perc_50(n) FROM x",
	}, "3")

	assertQueryPrepare(t, []string{
		"CREATE TABLE x (n INT)",
		"INSERT INTO x VALUES (5), (2), (4)",
		"SELECT perc(n, 75) FROM x",
	}, "5")
}
