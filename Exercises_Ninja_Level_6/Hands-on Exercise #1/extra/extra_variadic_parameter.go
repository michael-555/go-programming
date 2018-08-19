package main

import (
	"fmt"
)

func main() {
	sum(2, 3, 4, 5, 6, 7, 8, 9)

}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index pos,", i, "we are now adding,", v, "to the total which now is,", sum)
	}
	fmt.Println("The total is,", sum)
	return sum
}
