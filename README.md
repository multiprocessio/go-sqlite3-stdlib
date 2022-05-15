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

| Name(s) | Notes | Example |
| ======================== | ==== | === |
| repeat, replicate | |  `repeat('f', 5) = 'fffff'` |
| strpos, charindex | | `strpos('abc', 'b') = 1` |
| ltrim | Omit the second to default to trimming spaces | `ltrim('abccbad', 'abc') = 'd'` |
| rtrim | Omit the second to default to trimming spaces | `rtrim('abccbad', 'd') = 'abccbad'` |
| replace | | `replace('abc', 'c', 'd') = 'abd` |
| reverse | | `reverse('abc') = 'cba'` |
| lpad | Omit the second argument to default to padding with spaces | `lpad('22', 3, '0') = '022'` |
| rpad | Omit the second argument to default to padding with spaces | `rpad('22', 3, '0') = '220'`|

### Aggregation

| Name(s) | Notes | Example |
| ======================== | ==== | === |
| stddev, stdev, stddev_pop | | |
| mode | | |
| median | | |
| percentile, perc | Discrete | `perc(response_time, 95)` |
| percentile_25, perc_25, percentile_50, perc_50, percentile_75, perc_75, percentile_90, perc_90, percentile_95, perc_95, percentile_99, perc_99 | Discrete | `perc_99(response_time)` |
| percentile_cont, perc_cont | Continuous | `perc_cont(response_time, 95)` |
| percentile_cont_25, perc_cont_25, percentile_cont_50, perc_cont_50, percentile_cont_75, perc_cont_75, percentile_cont_90, perc_cont_90, percentile_cont_95, perc_cont_95, percentile_cont_99, perc_cont_99| Continuous | `perc_cont_99(response_cont_time)` |

### Net

| Name(s) | Notes | Example |
| ======================== | ==== | === |
| url_scheme | |  `url_scheme('https://x.com:90/home.html') = 'https'` |
* url_host | |  `url_host('https://x.com:90/home.html') = 'x.com:90'` |
* url_port | |  `url_port('https://x.com:90/home.html') = '90'` |
* url_path | | `url_path('https://x.com/some/path.html?p=123') = '/some/path.html'` | 
* url_param | | `url_param('https://x.com/home.html?p=123&z=%5B1%2C2%5D#section-1', 'z') = '[1,2]'` |
* url_fragment | | `url_fragment('https://x.com/home.html?p=123&z=%5B1%2C2%5D#section-1') = 'section-1'` |

### Date

Best effort family of date parsing and retrieval. Results will differ
depending on your computer's timezone.

| Name(s) | Notes | Example |
| ======================== | ==== | === |
| date_year | | `date_year('2021-04-05') = 2021` |
| date_month | January is 1, not 0 | `date_month('May 6, 2021') = 5` |
| date_day | | `date_day('May 6, 2021') = 6` |
| date_yearday | Day offset in year | `date_yearday('May 6, 2021') = 127` |
| date_hour | 24-hour | `date_hour('May 6, 2021 4:50 PM') = 16` |
| date_minute | | `date_minute('May 6, 2021 4:50') = 50` |
| date_second | | `date_second('May 6, 2021 4:50:20') = 20` |
| date_unix | | `date_unix('May 6, 2021 4:50:20') = 1588740620` |
| date_rfc3339 | | `date_unix('May 6, 2021 4:50:20') = 2020-05-06T04:50:20Z` |


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
