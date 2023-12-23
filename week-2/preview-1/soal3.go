package main

import (
	"fmt"
	"strconv"
)

func Cli() {
	var input string
	fmt.Println("masukkan angka :")
	fmt.Scanln(&input)

	num, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("input harus berupa angka")
	} else {
		for i := 1; i <= num; i++ {
			if i%2 == 0 && num%2 == 1 {
				fmt.Println(i)
			} else {
				fmt.Println(i)
			}
		}
	}
}
