package ch02

import (
	"fmt"
	"math"
)

func ch02Overflows() {
	i := 0
	for {
		fmt.Println(i)
		if i == math.MaxInt {
			fmt.Println("max reached")
			break
		}
		i = i + 1
	}
}
