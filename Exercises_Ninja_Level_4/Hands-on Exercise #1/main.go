package main

import (
	"fmt"
)

func main() {
	a := [5]int{25, 36, 75, 54, 247}

	for _, v := range a {
		fmt.Println(v)

	}
	fmt.Printf("The TYPE and lenght of the array is:%T", a)
}
