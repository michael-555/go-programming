package main

import (
	"fmt"
)

func main() {
	a := 30
	fmt.Printf("%d %b %#x\n", a, a, a) //Print decimal, binary & Hex

	// Shift the bits
	b := a << 1
	fmt.Printf("%d %b %#x", b, b, b)

}
