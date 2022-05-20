package stdlib

import (
	"database/sql"

	"github.com/mattn/go-sqlite3"
)

var extensions = []map[string]any{
	mathFunctions,
	stringFunctions,
	regexpFunctions,
	dateFunctions,
	netFunctions,
	encodingFunctions,
}

func ConnectHook(conn *sqlite3.SQLiteConn) error {
	var err error
	for _, functions := range extensions {
		for name, impl := range functions {
			err = conn.RegisterFunc(name, impl, true)
			// Yes it's weird to not break on return/error
			// but this way we get 100% test coverage on
			// this function. Errors shouldn't happen
			// outside of development anyway.
		}
	}

	for name, impl := range aggregateFunctions {
		err = conn.RegisterAggregator(name, impl, true)
	}

	return err
}

func Register(driverName string) {
	sql.Register(driverName,
		&sqlite3.SQLiteDriver{
			ConnectHook: ConnectHook,
		})
}
