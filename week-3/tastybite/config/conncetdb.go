package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func dbcon() string {
	name := "root"
	hostname := "127.0.0.1:3306"
	dbname := "latihandml_test"

	return fmt.Sprintf("%s:@tcp(%s)/%s", name, hostname, dbname)
}

func ConnectDb() *sql.DB {

	db, err := sql.Open("mysql", dbcon())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
