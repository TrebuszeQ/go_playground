package main

import "fmt"

func main() {
	animal := struct {
		name  string
		speak func() string
	}{
		name: "cat",
		speak: func() string {
			return "meow"
		},
	}

	fmt.Println(fmt.Printf("Our animal's name is %s and it says %s", animal.name, animal.speak()))
}
