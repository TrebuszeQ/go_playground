package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("GET /goodbye/", goodbyeHandler)
	mux.HandleFunc("GET /goodbye/{name}", goodbyeHandler)

	err := http.ListenAndServe(":8084", mux)
	if err != nil {
		panic(err)
	}
}
