package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
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

func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path

	if handlerFunc, ok := r.cachedHandlers[check]; ok {
		handlerFunc(res, req)
		return
	}

	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			r.cachedHandlers[check] = handlerFunc
			return
		}
	}
	http.NotFound(res, req)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbyeHandler(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := ""
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}

func main() {
	rr := newPathResolver()
	rr.Add("GET /hello", helloHandler)
	rr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*)?", goodbyeHandler)
	if err := http.ListenAndServe(":8080", rr); err != nil {
		panic(err)
	}

}
