package main

import (
	"fmt"
	"math/rand"
)

func Factorial(ch <-chan int) chan int {
	c := make(chan int)

	go func() {
		n := <-ch
		hasil := 1
		for i := 1; i <= n; i++ {
			hasil *= i

		}
		c <- hasil
		close(c)
	}()

	return c
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- rand.Intn(100)
		close(ch)

	}()

	result := <-Factorial(ch)
	fmt.Println(result)

}
