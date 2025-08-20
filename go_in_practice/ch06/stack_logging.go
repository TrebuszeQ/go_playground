package ch06

import (
	"runtime/debug"
)

func bar() {
	debug.PrintStack()
}

func foo() {
	bar()
}

func main() {
	foo()
}
