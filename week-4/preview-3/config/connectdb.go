package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	const (
		name     = "root"
		hostname = "127.0.0.1:3306"
		dbname   = "bookstore"
	)

	connection := fmt.Sprintf("%s:@tcp(%s)/%s", name, hostname, dbname)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
