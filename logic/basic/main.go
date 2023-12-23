package main

import "fmt"

func Gimme(array [3]int) {
	for i := 0; i < len(array); i++ {
		res := len(array)%2 == 1
		fmt.Println(res)
	}

}

func main() {
	arr := [3]int{6, 10, 20}
	Gimme()
}
