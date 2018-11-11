// +build mysql

package testfixtures

import (
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	databases = append(databases,
		databaseTest{
			name:       "mysql",
			connEnv:    "MYSQL_CONN_STRING",
			schemaFile: "testdata/schema/mysql.sql",
			helper:     &MySQL{},
		},
	)
}
