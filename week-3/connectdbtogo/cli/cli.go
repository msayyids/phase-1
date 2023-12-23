package cli

import (
	"connectdb/handler"
	"fmt"
)

type Cli struct {
	UserHandler *handler.UserHandler
}

func NewCli(uh *handler.UserHandler) *Cli {
	return &Cli{UserHandler: uh}
}

func (c *Cli) Show() {
	users, err := c.UserHandler.FindAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, user := range users {
		fmt.Printf("User ID: %d, Name: %s, Age: %d\n", user.Id, user.Name, user.Age)
	}
}
