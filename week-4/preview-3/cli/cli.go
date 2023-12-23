package cli

import (
	"bookstore/handler"
	"fmt"
	"os"
)

type Cli struct {
	DB *handler.Handler
}

func (c *Cli) Cli() {
	if len(os.Args) < 2 {
		fmt.Println("please scpecify a command : books,sales,customer,topauthor")
		return
	}

	command := os.Args[1]

	switch command {
	case "books":
		c.DB.ListBook()
	case "sales":
		c.DB.TotalSales()
	case "customer":
		c.DB.Coustomer()
	default:
		fmt.Println("unknow command:", command)
		fmt.Println("please scpecify a command : books,sales,customer,topauthor")
	}
}
