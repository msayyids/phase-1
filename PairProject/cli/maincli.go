package cli

import (
	"fmt"
	"log"
	"os"
	"pairproject/handler"
)

type Cli struct {
	CommandLi *handler.Handler
}

func (c *Cli) Cli() {
	fmt.Println("welcome to Games Store!")
	fmt.Println("[COMMAND]: 	Description")
	fmt.Println("[register]: 	Register/create new user")
	fmt.Println("[[login]: 	login into existing user")

	var inputCli string
	_, err := fmt.Scanln(&inputCli)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch inputCli {
	case "register":
		c.Register()
	case "login":
		c.Login()
	default:
		os.Exit(1)
	}

}
