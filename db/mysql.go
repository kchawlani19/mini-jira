// db/mysql.go
package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mini_jira")
	if err != nil {
		panic(err)
	}
	return db
}
