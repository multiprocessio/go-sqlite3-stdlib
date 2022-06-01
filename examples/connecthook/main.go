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
