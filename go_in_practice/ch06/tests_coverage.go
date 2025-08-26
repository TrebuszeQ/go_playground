package main

import (
	"testing"
)

func foo() int {
	return 1
}

func bar() int {
	return 2
}

func baz() int {
	return 3
}

func TestFoo(t *testing.T) {
	t.Run("testing foo", func(t *testing.T) {
		foo()
	})
}

func TestBar(t *testing.T) {
	t.Run("testing bar", func(t *testing.T) {
		bar()
	})
}

func main() {}
