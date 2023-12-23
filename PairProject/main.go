package main

import (
	"pairproject/cli"
	"pairproject/config"
	"pairproject/handler"
)

func main() {
	db := config.ConnectDb()
	h := &handler.Handler{DB: db}
	c := &cli.Cli{CommandLi: h}
	c.Cli()

}
