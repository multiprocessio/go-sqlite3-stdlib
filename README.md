# A standard library for mattn/go-sqlite3

As an alternative to compiling C extensions like
[extension-functions.c](https://www.sqlite.org/contrib) and
[sqlean](https://github.com/nalgeon/sqlean) into
[mattn/go-sqlite3](https://github.com/mattn/go-sqlite3), this package
implements many of these functions (and more from PostgreSQL) in Go.

This is mostly bindings to Go standard library functions or some
third-party libraries like [gonum](https://gonum.org/v1/gonum) and
[dateparse](https://github.com/araddon/dateparse).

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

# Functions

## Strings

| Name(s) | Notes | Example |
| ------------------------ | ---- | --- |
| repeat, replicate | |  `repeat('f', 5) = 'fffff'` |
| strpos, charindex | | `strpos('abc', 'b') = 1` |
| reverse | | `reverse('abc') = 'cba'` |
| lpad | Omit the second argument to default to padding with spaces | `lpad('22', 3, '0') = '022'` |
| rpad | Omit the second argument to default to padding with spaces | `rpad('22', 3, '0') = '220'`|

## Aggregation

| Name(s) | Notes | Example |
| ------------------------ | ---- | --- |
| stddev, stdev, stddev_pop | | |
| mode | | |
| median | | |
| percentile, perc | Discrete | `perc(response_time, 95)` |
| percentile_25, perc_25, percentile_50, perc_50, percentile_75, perc_75, percentile_90, perc_90, percentile_95, perc_95, percentile_99, perc_99 | Discrete | `perc_99(response_time)` |
| percentile_cont, perc_cont | Continuous | `perc_cont(response_time, 95)` |
| percentile_cont_25, perc_cont_25, percentile_cont_50, perc_cont_50, percentile_cont_75, perc_cont_75, percentile_cont_90, perc_cont_90, percentile_cont_95, perc_cont_95, percentile_cont_99, perc_cont_99| Continuous | `perc_cont_99(response_time)` |

## Net

| Name(s) | Notes | Example |
| ------------------------ | ---- | --- |
| url_scheme | |  `url_scheme('https://x.com:90/home.html') = 'https'` |
| url_host | | `url_host('https://x.com:90/home.html') = 'x.com:90'` |
| url_port | | `url_port('https://x.com:90/home.html') = '90'` |
| url_path | | `url_path('https://x.com/some/path.html?p=123') = '/some/path.html'` | 
| url_param | | `url_param('https://x.com/home.html?p=123&z=%5B1%2C2%5D#section-1', 'z') = '[1,2]'` |
| url_fragment | | `url_fragment('https://x.com/home.html?p=123&z=%5B1%2C2%5D#section-1') = 'section-1'` |

## Date

Best effort family of date parsing and retrieval. Results will differ
depending on your computer's timezone.

| Name(s) | Notes | Example |
| ------------------------ | ---- | --- |
| date_year | | `date_year('2021-04-05') = 2021` |
| date_month | January is 1, not 0 | `date_month('May 6, 2021') = 5` |
| date_day | | `date_day('May 6, 2021') = 6` |
| date_yearday | Day offset in year | `date_yearday('May 6, 2021') = 127` |
| date_hour | 24-hour | `date_hour('May 6, 2021 4:50 PM') = 16` |
| date_minute | | `date_minute('May 6, 2021 4:50') = 50` |
| date_second | | `date_second('May 6, 2021 4:50:20') = 20` |
| date_unix | | `date_unix('May 6, 2021 4:50:20') = 1588740620` |
| date_rfc3339 | | `date_rfc3339('May 6, 2021 4:50:20') = 2020-05-06T04:50:20Z` |

## Regexp

| Name(s) | Notes | Example |
| ------------------------ | ---- | --- |
| regexp | Go's regexp package, not PCRE. | `x REGEXP '[a-z]+$'`, `REGEXP('[a-z]+$', x)` |

## Math

| Name(s) | Notes | Example |
| ------------------------ | ---- | --- |
| acos | | `acos(n)` |
| acosh  | | `acosh(n)` |
| asin | |`asin(n)` |
| asinh  | | `asinh(n)` |
| atan | | `atan(n)` |
| atanh | | `atanh(n)` |
| ceil, ceiling | | `ceil(n)` |
| cos | | `ceil(n)` |
| cosh | | `cosh(n)` |
| degrees | | `degrees(radians)` |
| exp | e^n | `exp(n)` |
| floor | | `floor(n)` |
| ln, log | | `log(x)` |
| log10 | | `log10(x)` |
| log2 | | `log2(x)` |
| mod | | `mod(num, denom)` |
| pi | | `pi()` |
| pow, power | | `pow(base, exp)` |
| radians | | `radians(degrees)` |
| sin | | `sin(n)` |
| sinh | | `sinh(n)` |
| sqrt | | `sqrt(n)` |
| tan | | `tan(n)` |
| tanh | | `tanh(n)` |
| trunc, truncate | Rounds up to zero if negative, down to zero if positive. | `trunc(-10.9) = -10`, `trunc(10.4) = 10.0` |

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

