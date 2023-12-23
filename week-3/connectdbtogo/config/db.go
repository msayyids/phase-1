package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/nyobadb")
	if err != nil {
		log.Fatal((err.Error()))
	}


	err = db.Ping()
	if err != nil {
		log.Fatal((err.Error()))
	}

	fmt.Println("berhasil konek")
	return db

}
