package main

import "fmt"

// soal 1

func MergeSlices[T int | string](slice1 []T, slice2 []T) []T {
	//write code here
	var output []T
	output = append(slice1, slice2...)
	// output = append(output, slice2...)

	return output
}

func main() {
	//write code here
	intSlice1 := []int{1, 2, 3}
	intSlice2 := []int{4, 5, 6}
	mergedIntSlice := MergeSlices(intSlice1, intSlice2)
	fmt.Println(mergedIntSlice)
	//write response here
	strSlice1 := []string{"apple", "banana"}
	strSlice2 := []string{"orange", "grape"}
	mergedStrSlice := MergeSlices(strSlice1, strSlice2)
	fmt.Println(mergedStrSlice)

	// Soal 2

	foodProduct := Product[Food]{
		Name:  "Chips",
		Price: 10,
		Data:  Food{ExpiryDate: "2023-12-31"},
	}

	foodProduct.Getinfo()

	ElectronicsProduct := Product[Electronics]{
		Name:  "Chips",
		Price: 10,
		Data:  Electronics{ModelNumber: "abc123"},
	}

	ElectronicsProduct.Getinfo()

}
