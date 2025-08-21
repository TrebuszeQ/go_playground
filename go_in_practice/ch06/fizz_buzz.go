package ch06

import (
	"fmt"
	"strconv"
	"strings"
)

func fizzbuzz(n int64) string {
	var fizzbuzzes []string
	for i := int64(1); i <= n; i++ {
		v := ", "
		isThree := i%3 == 0
		isFive := i%5 == 0
		if isThree && isFive {
			v = "Fizz Buzz"
		} else if isThree {
			v = "Fizz"
		} else if isFive {
			v = "Buzz"
		} else {
			v = fmt.Sprintf("%d", i)
		}
		fizzbuzzes = append(fizzbuzzes, v)
	}

	return strings.Join(fizzbuzzes, " ")
}

func main() {
	var input string
	fmt.Println("Enter a number for fizzbuzz")
	fmt.Scanln(&input)

	numInput, err := strconv.ParseInt(input, 10, 16)
	if err != nil {
		panic("taht's not a number!")
	}

	result := fizzbuzz(numInput)
	fmt.Println("result:", result)
}
