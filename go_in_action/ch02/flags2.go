package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "World", "A name to say hello to.")

type language = string

var userLanguage language

validLanguages = []string{
	"en",
	"sp",
	"fr",
	"de"
}

func init() {
	flag.StringVar(&userLanguage, "lang", "en", "Language to use (en, sp, fr, de).")
	flag.Parse()
	if slices.Contains(validLanguages, userLanguage) == false {
		log.Fatalf("Invalid language %s. Use one of %v", userLanguage, validLanguages)
	}
}

func main() {
	switch userLanguage {
	case English:
		log.Printf("Hello %s!\n", *name)
	case Spanish:
		log.Printf("Hola %s!\n", *name)
	case French:
		log.Printf("Bonjour %s!\n", *name)
	case German:
		log.Printf("Hallo %s!\n", *name)
	}
}
