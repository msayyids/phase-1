package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbnames  = "FinalProjectP1"
)

func dbmysql(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func ConnectDb() *sql.DB {

	db, err := sql.Open("mysql", dbmysql(dbnames))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("konek")
	return db
}
