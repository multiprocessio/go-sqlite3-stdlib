package main

import (
	"database/sql"
	"fmt"

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
