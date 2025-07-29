package main

import (
	"net/http"
	"regexp"
)

type regexResolver struct {
	handlers       map[string]http.HandlerFunc
	cache          map[string]*regexp.Regexp
	cachedHandlers map[string]http.HandlerFunc
}

func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers:       make(map[string]http.HandlerFunc),
		cache:          make(map[string]*regexp.Regexp),
		cachedHandlers: make(map[string]http.HandlerFunc),
	}
}

func main() {

}
