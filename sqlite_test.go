// +build sqlite

package testfixtures

import (
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	databases = append(databases, databaseTest{
		name:       "sqlite3",
		connEnv:    "SQLITE_CONN_STRING",
		schemaFile: "testdata/schema/sqlite.sql",
		helper:     &SQLite{},
	})
}
