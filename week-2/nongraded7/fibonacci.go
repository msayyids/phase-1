package main

import (
	"fmt"
	"sync"
)

func fibonacci(n int, c chan<- any, wg *sync.WaitGroup) {
	defer wg.Done()
	if n <= 0 {
		c <- 0
	} else if n == 1 {
		c <- 1
	}

	prev, current := 0, 1
	for i := 2; i <= n; i++ {
		prev, current = current, prev+current
	}

	c <- current
}

func soal1() {
	var wg sync.WaitGroup
	var input int
	fmt.Println("masukkan angka:")
	fmt.Scanln(&input)
	c := make(chan any, input)
	wg.Add(input)
	go fibonacci(input, c, &wg)

	go func() {
		wg.Wait()
		close(c)
	}()

	result := <-c

	fmt.Println(result)
}
