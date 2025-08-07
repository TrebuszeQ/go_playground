package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrTimeout = errors.New("the request timed out")

func SendRequest(req string) (string, error) {
	return "", fmt.Errorf("we got an error: %w ", ErrTimeout)
}

func main() {
	if _, err := SendRequest("Hello "); err != nil {
		if errors.Is(err, ErrTimeout) {
			log.Println("we got a timeout error")
		} else {
			log.Println("we got some other error")
		}
	}
}
