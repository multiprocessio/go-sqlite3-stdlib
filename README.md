# A standard library for mattn/go-sqlite3

As an alternative to compiling C extensions like
[extension-functions.c](https://www.sqlite.org/contrib) and
[sqlean](https://github.com/nalgeon/sqlean) into
[mattn/go-sqlite3](https://github.com/mattn/go-sqlite3), this package
implements many of these functions (and more from PostgreSQL) in Go.

These are in addition to [all builtin
functions](https://www.sqlite.org/lang_corefunc.html) provided by
SQLite.

Continue reading for all functions, notes and examples.

# Why would I use this?

This library is used in
[DataStation](https://github.com/multiprocessio/datastation) and
[dsq](https://github.com/multiprocessio/dsq) to simplify and power
data analysis in SQL.

![Analyzing logs with SQL in DataStation](./screenshot.png)

Read the [DataStation blog
post](https://datastation.multiprocess.io/docs/0.11.0-release-notes.html)
to better understand the background.

# Example

```go
package main

import (
	"fmt"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	stdlib "github.com/multiprocessio/go-sqlite3-stdlib"
)

func main() {
	stdlib.Register("sqlite3_ext")
	db, err := sql.Open("sqlite3_ext", ":memory:")
	if err != nil {
		panic(err)
	}

	var s string
	err = db.QueryRow("SELECT repeat('x', 2)").Scan(&s)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
```

Alternatively if you want to be able to add your own additional
extensions you can just use the `ConnectHook`:

```go
package main

import (
	"database/sql"
	"fmt"

	sqlite3 "github.com/mattn/go-sqlite3"
	stdlib "github.com/multiprocessio/go-sqlite3-stdlib"
)

func main() {
	sql.Register("sqlite3_ext",
		&sqlite3.SQLiteDriver{
			ConnectHook: stdlib.ConnectHook,
		})
	db, err := sql.Open("sqlite3_ext", ":memory:")
	if err != nil {
		panic(err)
	}

	var s string
	err = db.QueryRow("SELECT repeat('x', 2)").Scan(&s)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
```

# Functions

## Strings

| Name(s)           | Notes                                                     | Example                                                        |
| ----------------- | --------------------------------------------------------- | -------------------------------------------------------------- |
| repeat, replicate |                                                           | `repeat('f', 5) = 'fffff'`                                     |
| strpos, charindex | 0-indexed position of substring in string                 | `strpos('abc', 'b') = 1`                                       |
| reverse           |                                                           | `reverse('abc') = 'cba'`                                       |
| lpad              | Omit the third argument to default to padding with spaces | `lpad('22', 3, '0') = '022'`                                   |
| rpad              | Omit the third argument to default to padding with spaces | `rpad('22', 3, '0') = '220'`                                   |
| len               | Shorthand for `length`                                    | `len('my string') = '9'`                                       |
| split_part        | Split string an take nth split piece                      | `split('1,2,3', ',', 0) = '1'`, `split('1,2,3', ',' -1) = '3'` |
| regexp            | Go's regexp package, not PCRE                             | `x REGEXP '[a-z]+$'`, `REGEXP('[a-z]+$', x)`                   |
| regexp_count      | Number of times the regexp matches in string              | `regexp_count('abc1', '[a-z]1') = '1'`                         |
| regexp_split_part | Regexp equivalent of `split_part`                         | `regexp_split_part('ab12', '[a-z]1', 0) = 'a'`                 |

## Aggregation

Most of these are implemented as bindings to
[gonum](https://gonum.org/v1/gonum).

| Name(s)                                                                                                                                                                                                    | Notes      | Example                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- | ------------------------------ |
| stddev, stdev, stddev_pop                                                                                                                                                                                  |            | `stddev(n)`                    |
| mode                                                                                                                                                                                                       |            | `mode(n)`                      |
| median                                                                                                                                                                                                     |            | `median(n)`                    |
| percentile, perc                                                                                                                                                                                           | Discrete   | `perc(response_time, 95)`      |
| percentile_25, perc_25, percentile_50, perc_50, percentile_75, perc_75, percentile_90, perc_90, percentile_95, perc_95, percentile_99, perc_99                                                             | Discrete   | `perc_99(response_time)`       |
| percentile_cont, perc_cont                                                                                                                                                                                 | Continuous | `perc_cont(response_time, 95)` |
| percentile_cont_25, perc_cont_25, percentile_cont_50, perc_cont_50, percentile_cont_75, perc_cont_75, percentile_cont_90, perc_cont_90, percentile_cont_95, perc_cont_95, percentile_cont_99, perc_cont_99 | Continuous | `perc_cont_99(response_time)`  |

## Net

| Name(s)      | Notes | Example                                                                               |
| ------------ | ----- | ------------------------------------------------------------------------------------- |
| url_scheme   |       | `url_scheme('https://x.com:90/home.html') = 'https'`                                  |
| url_host     |       | `url_host('https://x.com:90/home.html') = 'x.com:90'`                                 |
| url_port     |       | `url_port('https://x.com:90/home.html') = '90'`                                       |
| url_path     |       | `url_path('https://x.com/some/path.html?p=123') = '/some/path.html'`                  |
| url_param    |       | `url_param('https://x.com/home.html?p=123&z=%5B1%2C2%5D#section-1', 'z') = '[1,2]'`   |
| url_fragment |       | `url_fragment('https://x.com/home.html?p=123&z=%5B1%2C2%5D#section-1') = 'section-1'` |

## Date

Best effort family of date parsing (uses
[dateparse](https://github.com/araddon/dateparse)) and date part
retrieval. Results will differ depending on your computer's timezone.

| Name(s)      | Notes               | Example                                                      |
| ------------ | ------------------- | ------------------------------------------------------------ |
| date_year    |                     | `date_year('2021-04-05') = 2021`                             |
| date_month   | January is 1, not 0 | `date_month('May 6, 2021') = 5`                              |
| date_day     |                     | `date_day('May 6, 2021') = 6`                                |
| date_yearday | Day offset in year  | `date_yearday('May 6, 2021') = 127`                          |
| date_hour    | 24-hour             | `date_hour('May 6, 2021 4:50 PM') = 16`                      |
| date_minute  |                     | `date_minute('May 6, 2021 4:50') = 50`                       |
| date_second  |                     | `date_second('May 6, 2021 4:50:20') = 20`                    |
| date_unix    |                     | `date_unix('May 6, 2021 4:50:20') = 1588740620`              |
| date_rfc3339 |                     | `date_rfc3339('May 6, 2021 4:50:20') = 2020-05-06T04:50:20Z` |

## Math

| Name(s)         | Notes                                                    | Example                                    |
| --------------- | -------------------------------------------------------- | ------------------------------------------ |
| acos            |                                                          | `acos(n)`                                  |
| acosh           |                                                          | `acosh(n)`                                 |
| asin            |                                                          | `asin(n)`                                  |
| asinh           |                                                          | `asinh(n)`                                 |
| atan            |                                                          | `atan(n)`                                  |
| atanh           |                                                          | `atanh(n)`                                 |
| ceil, ceiling   |                                                          | `ceil(n)`                                  |
| cos             |                                                          | `ceil(n)`                                  |
| cosh            |                                                          | `cosh(n)`                                  |
| degrees         |                                                          | `degrees(radians)`                         |
| exp             | e^n                                                      | `exp(n)`                                   |
| floor           |                                                          | `floor(n)`                                 |
| ln, log         |                                                          | `log(x)`                                   |
| log10           |                                                          | `log10(x)`                                 |
| log2            |                                                          | `log2(x)`                                  |
| mod             |                                                          | `mod(num, denom)`                          |
| pi              |                                                          | `pi()`                                     |
| pow, power      |                                                          | `pow(base, exp)`                           |
| radians         |                                                          | `radians(degrees)`                         |
| sin             |                                                          | `sin(n)`                                   |
| sinh            |                                                          | `sinh(n)`                                  |
| sqrt            |                                                          | `sqrt(n)`                                  |
| tan             |                                                          | `tan(n)`                                   |
| tanh            |                                                          | `tanh(n)`                                  |
| trunc, truncate | Rounds up to zero if negative, down to zero if positive. | `trunc(-10.9) = -10`, `trunc(10.4) = 10.0` |

## Encoding

| Name(s)     | Notes                         | Example          |
| ----------- | ----------------------------- | ---------------- |
| base64      | Convert string to base64      | `base64(s)`      |
| from_base64 | Convert string from base64    | `from_base64(s)` |
| base32      | Convert string to base32      | `base32(s)`      |
| from_base32 | Convert string from base32    | `from_base32(s)` |
| md5         | Hex md5 sum of string         | `md5(s)`         |
| sha1        | Hex sha1 sum of string        | `sha1(s)`        |
| sha256      | Hex sha256 sum of string      | `sha256(s)`      |
| sha512      | Hex sha512 sum of string      | `sha512(s)`      |
| sha3_256    | Hex sha3_256 sum of string    | `sha3_256(s)`    |
| sha3_512    | Hex sha3_512 sum of string    | `sha3_512(s)`    |
| blake2b_256 | Hex blake2b_256 sum of string | `blake2b_256(s)` |
| blake2b_512 | Hex blake2b_512 sum of string | `blake2b_512(s)` |

# How is this tested?

There is 95% test coverage and automated tests on Windows, macOS and
Linux.

# I just want to use it as a CLI or GUI

See [dsq](https://github.com/multiprocessio/dsq) (a command-line tool
for executing SQL on data files) and
[DataStation](https://github.com/multiprocessio/datastation), a GUI
application for querying and building reports with data from
databases, servers, and files.

# Contribute

Join the [#dev channel on the Multiprocess Labs
Discord](https://discord.gg/22ZCpaS9qm).

If you have an idea for a new function, say so on the Discord channel
or open an issue here.

Make sure the function doesn't already exist in dsq (or the sqlite3
CLI).

# License

This software is licensed under an Apache 2.0 license.
