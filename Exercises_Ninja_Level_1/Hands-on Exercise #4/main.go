package main

import (
	"fmt"
)

type deposit int

var x deposit

func main() {
	fmt.Println(x)                    // Zero Value
	fmt.Println(fmt.Sprintf("%T", x)) // print the type
	x = 42
	fmt.Println(x)
}
