// +build oracle

package testfixtures

import (
	_ "github.com/mattn/go-oci8"
)

func init() {
	databases = append(databases,
		databaseTest{
			name:       "oci8",
			connEnv:    "ORACLE_CONN_STRING",
			schemaFile: "testdata/schema/oracle.sql",
			helper:     &Oracle{},
		},
	)
}
