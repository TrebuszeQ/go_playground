package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const MAX_TIMEOUTS = 5

var ErrTimeout = errors.New("The request timed out")
var ErrRejected = errors.New("The request was rejected")

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func SendRequest(req string) (string, error) {
	switch rand.Intn(3) % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}

func main() {
	response, err := SendRequest("Hello")

	if errors.Is(err, ErrTimeout) {
		timeouts := 0
		for err == errors.Is(err, ErrTimeout) {
			timeouts++
			fmt.Println("Timeout. Retrying.")

			if timeouts == MAX_TIMEOUTS {
				panic("too many timeouts!")
			}
			response, err = SendRequest("Hello")
		}
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
