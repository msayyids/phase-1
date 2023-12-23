package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB {
	const (
		username = "root"
		Password = ""
		hostname = "127.0.0.1:3306"
		dbname   = "nyobadb"
	)

	dbSource := fmt.Sprintf("%s:@tcp(%s)/%s", username, hostname, dbname)
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("berhasil konek")
	return db

}
