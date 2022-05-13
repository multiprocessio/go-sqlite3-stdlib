package stdlib

import (
	"database/sql"

	"github.com/mattn/go-sqlite3"
)

var extensions = []map[string]any{
	mathFunctions,
	stringFunctions,
	regexpFunctions,
}

func Register(driverName string) {
	sql.Register(driverName,
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				for _, functions := range extensions {
					for name, impl := range functions {
						err := conn.RegisterFunc(name, impl, true)
						if err != nil {
							return err
						}
					}
				}

				for name, impl := range aggregateFunctions {
					err := conn.RegisterAggregator(name, impl, true)
					if err != nil {
						return err
					}
				}

				return nil
			},
		})
}
