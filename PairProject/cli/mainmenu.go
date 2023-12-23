package cli

import (
	"fmt"
	"log"
)

func (c *Cli) MainMenu() {
	fmt.Println("\nwelcome to Games Store")
	fmt.Println("[COMMAND]: 		Description")
	fmt.Println("[buy-games]: 		buy games")
	fmt.Println("[list-whistlist]: 	Retrieve all Whistlist")
	fmt.Println("[my-games]: 		see my games")
	fmt.Println("or press anything to back to main menu")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch input {
	case "buy-games":
		c.CommandLi.BuyGames(email)
	case "list-whistlist":
		c.CommandLi.Whistlist(email)
	default:
		c.MainMenu()
	}
}
