package main

import (
	"fmt"
)

type Animal interface {
	speak()
}

type Cat struct{}

func (c Cat) speak() {
	fmt.Println("meow")
}

func NewCat() *Cat {
	return &Cat{}
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("woof")
}

func NewDog() *Dog {
	return &Dog{}
}

type Llama struct{}

func (l Llama) speak() {
	fmt.Println("Nondescript animal noise?")
}

func NewLlama() *Llama {
	return &Llama{}
}

func main() {
	var a Animal
	c := NewCat()
	a = c
	a.speak()

	d := NewDog()
	a = d
	a.speak()

	l := NewLlama()
	a = l
	l.speak()
}
