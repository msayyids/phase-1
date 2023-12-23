package main

import "fmt"

func jumlah(arr []float64) {
	var total float64
	for i := 0; i < len(arr)-1; i++ {
		total += arr[i]

	}
	fmt.Println(total)
}

func rataRata(arr []float64) {
	var rataRata float64
	var totalArray float64
	jumahIndex := len(arr)

	for i := 0; i < len(arr); i++ {
		totalArray += arr[i]
		rataRata = totalArray / float64(jumahIndex)
	}

	fmt.Println(rataRata)
}

func median(arr []float64) float64 {
	index := len(arr)
	if index == 0 {
		return 0
	}

	mid := index / 2
	if index%2 == 1 {
		return arr[mid]
	} else {
		return (arr[mid-1] + arr[mid]) / 2
	}
}

func Looping2() {
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}
	fmt.Println("output jumlah")
	jumlah(slice1)
	jumlah(slice2)
	fmt.Println("output rata rata")
	rataRata(slice1)
	rataRata(slice2)
	fmt.Println("output median")
	fmt.Println(median(slice1))
	fmt.Println(median(slice2))
}
