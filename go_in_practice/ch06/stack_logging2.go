package ch06

import (
	"fmt"
	"runtime"
)

func bar() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	fmt.Printf("Trace:\n %s\n", buf)
}
