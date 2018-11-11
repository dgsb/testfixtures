// +build sqlserver

package testfixtures

import (
	_ "github.com/denisenkom/go-mssqldb"
)

func init() {
	databases = append(databases,
		databaseTest{
			name:       "mssql",
			connEnv:    "SQLSERVER_CONN_STRING",
			schemaFile: "testdata/schema/sqlserver.sql",
			helper:     &SQLServer{},
		},
	)
}
