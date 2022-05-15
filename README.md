# A standard library for mattn/go-sqlite3

As an alternative to compiling C extensions like
[extension-functions.c](https://www.sqlite.org/contrib) and
[sqlean](https://github.com/nalgeon/sqlean) into
[mattn/go-sqlite3](https://github.com/mattn/go-sqlite3), this package
implements many of these functions (and more from PostgreSQL) in Go.

## Example

```go
package main

import (
  "database/sql"
  
  stdlib "github.com/multiprocessio/go-sqlite3-stdlib"
)

func main() {
  stdlib.Register("sqlite3_ext")
  db, err := sql.Open("sqlite3_ext", ":memory:")
  if err != nil {
    panic(err)
  }
  
  var s string
  err = db.QueryRow("SELECT repeat('x', 2)").Scan(%s)
  if err != nil {
    panic(err)
  }
  
  // s == "xx"
}
```

## Functions

### Regexp

* regexp: `x REGEXP '[a-z]+$'` (uses Go's regexp, not PCRE)

### Strings

* repeat, replicate: `repeat(string, ntimes)`
* strpos, charindex: `strpos(needle, haystack)`
* trim, ltrim: `ltrim('abccbad', 'abc') = 'd'`
* rtrim: `rtrim('abccbad', 'd') = 'abccbad'`
* replace: `replace(string, what, with)`
* reverse: `reverse(string)`
* lpad: `lpad(string, length[, what])`
* rpad: `rpad(string, length[, what])`

### Math

* acos
* acosh 
* asin
* asinh 
* atan
* atanh
* ceil
* ceiling
* cos
* cosh
* degrees
* exp
* floor
* ln, log
* log10
* log2
* mod
* pi
* pow, power
* radians
* sin
* sinh
* sqrt
* tan
* tanh
* trunc, truncate

### Aggregation

* stddev, stdev, stddev_pop
* mode
* median
* percentile, perc: `perc(response_time, 95)`; discrete
* percentile_25, perc_25, percentile_50, perc_50, percentile_75, perc_75, percentile_90, perc_90, percentile_95, perc_95, percentile_99, perc_99: `perc_99(response_time)`; discrete
* percentile_cont, perc_cont: `perc_cont(response_time, 95)`; continuous
* percentile_cont_25, perc_cont_25, percentile_cont_50, perc_cont_50, percentile_cont_75, perc_cont_75, percentile_cont_90, perc_cont_90, percentile_cont_95, perc_cont_95, percentile_cont_99, perc_cont_99: `perc_cont_99(response_cont_time)`; continuous

### Net

* url_scheme: `url_scheme('https://x.com:90/some/path.html?p=123&z=%5B1%2C2%5D#section-1') = 'https'`
* url_host: `url_host('https://x.com:90/some/path.html?p=123&z=%5B1%2C2%5D#section-1') = 'x.com:90'`
* url_port: `url_port('https://x.com:90/some/path.html?p=123&z=%5B1%2C2%5D#section-1') = '90'`
* url_path: `url_path('https://x.com:90/some/path.html?p=123&z=%5B1%2C2%5D#section-1') = '/some/path.html'`
* url_param: `url_param('https://x.com:90/some/path.html?p=123&z=%5B1%2C2%5D#section-1', 'z') = '[1,2]'`
* url_fragment: `url_fragment('https://x.com:90/some/path.html?p=123&z=%5B1%2C2%5D#section-1') = 'section-1'`

### Date

Best effort family of date parsing and retrieval:

* date_year: best effort date parsing, returns year
* date_month
* date_day: day in month
* date_yearday: day in year
* date_hours: 24 hour
* date_minutes
* date_seconds
* date_millis: offset from seconds, not aboslute
* date_unix: best effort convert date to unix timestamp
