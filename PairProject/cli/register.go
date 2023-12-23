package cli

import (
	"fmt"
	"log"
)

func (c *Cli) Register() {
	var name, email, password string
	var age int

	fmt.Println("Please input name:")
	_, errName := fmt.Scanln(&name)

	fmt.Println("Please input email:")
	_, errEmail := fmt.Scanln(&email)

	fmt.Println("Please input password:")
	_, errPassword := fmt.Scanln(&password)

	fmt.Println("Please input age:")
	_, errAge := fmt.Scanln(&age)

	err := c.CommandLi.Register(name, email, password, age)
	if err != nil {
		fmt.Println("failed to sign up...")
	} else {
		fmt.Println("\nSuccess register user!")
	}

	if errName != nil ||
		errEmail != nil ||
		errPassword != nil ||
		errAge != nil {
		log.Fatal("failed to scan input")
	}

	c.Cli()

}
