package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type german struct{}

func (g german) LanguageName() string {
	return fmt.Printf(" i can speak ")
}

func main() {

}
