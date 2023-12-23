package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
	age  int
	job  string
}

func (p *Person) Getinfo(wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	c <- fmt.Sprintf("nama saya %s , umur saya %d tahun pekerjaan saya %s", p.name, p.age, p.job)

}

func (p *Person) AddYear() {
	p.age++

	if p.age >= 50 {
		p.job = "pensiun"
	}

}

func soal1() {

	var wg sync.WaitGroup

	var person = []*Person{
		{name: "ali", age: 45, job: "frontend"},
		{name: "agus", age: 45, job: "backend"},
		{name: "ari", age: 45, job: "ui ux"},
	}
	c := make(chan string, len(person))

	wg.Add(len(person))

	for i := 0; i < 5; i++ {
		for _, persons := range person {
			persons.AddYear()
		}
	}

	for _, persons := range person {
		go persons.Getinfo(&wg, c)

	}
	wg.Wait()

	close(c)

	for result := range c {
		fmt.Println(result)
	}

}
