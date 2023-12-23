package handler

import (
	"context"
	"fmt"
	"log"
	"tastybite/entity"
)

func (u *UserHandler) AddMenuItems() {
	var addmenu entity.MenuList

	inputAndScan := func(input interface{}, prompt string) {
		fmt.Printf("input %s\n", prompt)
		_, err := fmt.Scanln(input)
		if err != nil {
			fmt.Println("Invalid input.")
		}
	}

	inputAndScan(&addmenu.Name, "nama menu")
	inputAndScan(&addmenu.Desc, "desc")
	inputAndScan(&addmenu.Price, "price")
	inputAndScan(&addmenu.Category, "category")

	query := `insert into MenuItems (name,description,price,category)
	values
		(?,?,?,?);`
	ctx := context.Background()
	_, err := u.DB.ExecContext(ctx, query, addmenu.Name, addmenu.Desc, addmenu.Price, addmenu.Category)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("sukses insert menu items")
		u.ShowMenu()
	}

}

func (u *UserHandler) ShowMenu() {
	query := `select * from MenuItems;`

	ctx := context.Background()

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-20s %-20s %-20s %-20s %-20s\n", "item id", "name", "desc", "price", "category")
	fmt.Printf("%-20s %-20s %-20s %-20s %-20s\n", "-------", "-------", "-------", "-------", "-------")

	for rows.Next() {
		var showMenu entity.MenuList

		err = rows.Scan(&showMenu.ItemId, &showMenu.Name, &showMenu.Desc, &showMenu.Price, &showMenu.Category)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-20d %-20s %-20s $%-20v %-20s\n", showMenu.ItemId, showMenu.Name, showMenu.Desc, showMenu.Price, showMenu.Category)
	}
}
