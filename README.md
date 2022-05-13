# A standard library for mattn/go-sqlite3

As an alternative to compiling C extensions like extension-functions.c
and sqlean into mattn/go-sqlite3, this package implements many of
these functions (and more from PostgreSQL) in Go.

## Example

```
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

* regexp

### Strings

* repeat
* replicate
* strpos
* charindex
* ltrim
* rtrim
* trim
* replace
* reverse
* lpad
* rpad

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
* mode (for integers), modestr (for strings)
