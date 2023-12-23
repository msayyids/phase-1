package main

import (
	"fmt"
	"strconv"
)

type Angka struct {
	Val int
}

type Calculator interface {
	Calc() int
	Cli() int
}

func (a *Angka) Calc() int {
	data := []Angka{}
	var input string
	var opr string
	for {
		fmt.Println("masukkan angka :")
		fmt.Scanln(&input)

		if input == "=" {
			break
		}

		fmt.Println("+ - * / :")
		fmt.Scanln(&opr)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input harus berupa angka")
			continue
		}
		data = append(data, Angka{Val: num})
	}
	hasil := 0
	// for _, v := range data {

	// }

	return hasil

}
