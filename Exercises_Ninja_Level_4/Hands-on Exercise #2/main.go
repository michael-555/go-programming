package main

import (
	"fmt"
)

func main() {
	a := []int{25, 36, 75, 54, 247, 42, 75, 45, 12, 74}

	for _, v := range a {
		fmt.Println(v)

	}
	fmt.Printf("The TYPE and lenght of the slice is: %T and %d", a, len(a))
}
