package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDB() *sql.DB {

	dsn := "root:root@tcp(localhost:3306)/mini_jira"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("MySQL connected")
	return db
}
