package ch02

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("Custom error message")
	}

	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a: %d, b: %d, userId: %d", a, b, os.Geteuid())
	}
	return nil
}

func ch02Check() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("check() executed fine")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)
	if err.Error() == "Custom error message" {
		fmt.Println("Got custom error")
	}

	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("int value is", i)
	}

	i, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	}
}
