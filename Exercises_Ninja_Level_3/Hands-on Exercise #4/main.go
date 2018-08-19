package main

import (
	"fmt"
)

func main() {
	b := 1974
	for {
		if b > 2018 {
			break
		}
		fmt.Println(b)
		b++
	}
}
