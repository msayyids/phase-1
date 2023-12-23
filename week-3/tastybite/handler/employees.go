package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tastybite/entity"
)

type UserHandler struct {
	DB *sql.DB
}

func (u *UserHandler) AddEmployees(firstName, lastName, position string) error {

	if firstName == "" || lastName == "" || position == "" {
		return errors.New("jangan dikosongin kosong")
	}

	query := `insert into employees (first_name,last_name,position) 
	values (?, ?, ?);`
	ctx := context.Background()
	_, err := u.DB.ExecContext(ctx, query, firstName, lastName, position)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (u *UserHandler) ShowEmployees() {

	query := `select * from employees`

	ctx := context.Background()

	rows, err := u.DB.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-20s %-20s %-20s %-20s\n", "employee id", "first name", "last name", "position")
	fmt.Printf("%-20s %-20s %-20s %-20s\n", "-------", "-------", "-------", "-------")

	for rows.Next() {
		var employees entity.Employees
		err = rows.Scan(&employees.Employee_id, &employees.First_name, &employees.Last_name, &employees.Position)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-20d %-20s %-20s %-20s\n", employees.Employee_id, employees.First_name, employees.Last_name, employees.Position)

	}
}
