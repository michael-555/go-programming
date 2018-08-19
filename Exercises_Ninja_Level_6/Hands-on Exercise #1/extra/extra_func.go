package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Mike")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Flemming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from foo")
}

// with one arg (PASS BY VALUE)
func bar(s string) {
	fmt.Println("Hello,", s)
}

// Return value
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}
