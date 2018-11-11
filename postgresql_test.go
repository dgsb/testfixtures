// +build postgresql

package testfixtures

import (
	_ "github.com/lib/pq"
)

func init() {
	databases = append(databases,
		databaseTest{
			name:       "postgres",
			connEnv:    "PG_CONN_STRING",
			schemaFile: "testdata/schema/postgresql.sql",
			helper:     &PostgreSQL{},
			subtest:    "",
		},
		databaseTest{
			name:       "postgres",
			connEnv:    "PG_CONN_STRING",
			schemaFile: "testdata/schema/postgresql.sql",
			helper:     &PostgreSQL{UseAlterConstraint: true},
			subtest:    "postgres with alter constraint",
		},
	)
}
