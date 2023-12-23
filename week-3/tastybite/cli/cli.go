package cli

import (
	"fmt"
	"tastybite/entity"
	"tastybite/handler"
)

type CLi struct {
	DB *handler.UserHandler
}

func (c *CLi) Cli() {
	for {
		var input int
		fmt.Println("choose :\n1.Add employees\n2.show employees\n3.add menu\n4.show menu\n9.exit")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var employees entity.Employees

			fmt.Println("input first name")
			fmt.Scanln(&employees.First_name)
			fmt.Println("input last name")
			fmt.Scanln(&employees.Last_name)
			fmt.Println("input position")
			fmt.Scanln(&employees.Position)

			err := c.DB.AddEmployees(employees.First_name, employees.Last_name, employees.Position)
			if err != nil {
				panic(err)
			}
		case 2:
			c.DB.ShowEmployees()
		default:
			break
		}
	}

}
