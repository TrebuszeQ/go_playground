package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name string `help:"the name or type of any animal, as long as it is a cat or dog`
}

func (a Animal) speak() string {
	switch a.name {
	case "cat":
		return "meow"
	case "dog":
		return "woof"
	default:
		if member, ok := reflect.TypeOf(a).FieldByName("name"); ok {
			return fmt.Sprintf("Invalid animal name: %s",
				member.Tag.Get("help"))
		}
		return "Nondescript animal noise?"
	}
}

func main() {

}
