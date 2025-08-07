package main

import (
	"fmt"
	"log"
)

type Animal struct {
	string
}

func (a Animal) speak() {
	log.Println(a.string)
}

func main() {
	a := Animal{
		"cat",
	}
	fmt.Println(a.speak())
	a.string = "dog"
	fmt.Println(a.speak())
}
