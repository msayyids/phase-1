package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func readFile(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	file := os.Args[1]
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		c <- fileScanner.Text()
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan string)

	wg.Add(1)

	go readFile(c, &wg)

	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}
