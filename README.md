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
* lpad: `lpad(string, length, what = ' ')`
* rpad: `rpad(string, length, what = ' ')`

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

### Aggregate

* stddev, stdev, stddev_pop
* mode
* median
* range
