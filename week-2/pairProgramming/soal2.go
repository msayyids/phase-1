package main

import "fmt"

type Product[T any] struct {
	Name  string
	Price int
	Data  T
}

type Food struct {
	ExpiryDate string
}

type Electronics struct {
	ModelNumber string
}

func (p Product[T]) Getinfo() {
	fmt.Printf("Product Name : %v\n", p.Name)
	fmt.Printf("Product Price : %v\n", p.Price)
	fmt.Printf("Product data : %v\n", p.Data)
}

// foodProduct := GetInfo[Food]{
// 	Name:  "Chips",
// 	Price: 10,
// 	Data:  Food{ExpiryDate: "2023-12-31"},
//    }
