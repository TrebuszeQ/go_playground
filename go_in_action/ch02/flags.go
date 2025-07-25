package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")
var inSpanish bool

func init() {
	flag.BoolVar(&inSpanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&inSpanish, "s", false, "Use Spanish language.")
	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t-%s: %s (Default: '%s)\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	})
	flag.Parse()
}

func main() {
	if inSpanish {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
