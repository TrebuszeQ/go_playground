package main

import (
	"fmt"
)

func getStringsNamed() (first string, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func main() {
	n1, n2 := getStringsNamed()
	fmt.Println(n1, n2)
}
