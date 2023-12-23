package cli

import (
	"fmt"
	"log"
)

var email, password string

func (c *Cli) Login() {
	fmt.Println("\nPlease Login!")

	fmt.Println("Please input email:")
	_, erremail := fmt.Scanln(&email)

	fmt.Println("Please input password:")
	_, errPassword := fmt.Scanln(&password)

	if erremail != nil ||
		errPassword != nil {
		log.Fatal("failed to scan input")
	}

	_, err := c.CommandLi.Login(email, password)
	if err != nil {

		fmt.Println("\nFailed to login!")
		c.Login()
	} else {
		fmt.Println("\nSuccess login user!")
		c.MainMenu()
	}

}
