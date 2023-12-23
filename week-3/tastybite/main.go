package main

import (
	"tastybite/cli"
	"tastybite/config"
	"tastybite/handler"
)

func main() {
	db := config.ConnectDb()
	uh := &handler.UserHandler{DB: db}
	cli := &cli.CLi{DB: uh}
	cli.Cli()
}
