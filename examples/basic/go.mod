module basic

go 1.18

replace github.com/mattn/go-sqlite3 v1.14.13 => github.com/multiprocessio/go-sqlite3 v1.14.14-0.20220513213203-12637a65d5d7

require (
	github.com/mattn/go-sqlite3 v1.14.13
	github.com/multiprocessio/go-sqlite3-stdlib v0.0.0-20220515032354-f4fe715d8da1
)

require (
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de // indirect
	gonum.org/v1/gonum v0.11.0 // indirect
)
