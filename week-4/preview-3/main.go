package main

import (
	"bookstore/cli"
	"bookstore/config"
	"bookstore/handler"
)

func main() {
	db := config.ConnectDb()
	h := &handler.Handler{DB: db}
	cli := &cli.Cli{DB: h}
	cli.Cli()
}
