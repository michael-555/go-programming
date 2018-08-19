package main

import (
	"fmt"
)

const a string = "Google"
const b = 1.2

func main() {
	fmt.Printf("TYPED const: %v %T\nUNTYPED const: %v %T", a, a, b, b)

}
