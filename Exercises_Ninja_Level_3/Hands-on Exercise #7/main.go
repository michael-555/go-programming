package main

import (
	"fmt"
)

func main() {
	x := 55.9
	y := fmt.Sprintf("%T", x)
	if y == "int" {
		fmt.Println("its a int")
	} else if y == "string" {
		fmt.Println("its a string")
	} else {
		fmt.Println("its a", y)
	}
}
