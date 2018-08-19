package main

import (
	"fmt"
)

const (
	a = 2018 + iota
	b
	c
	d
)

func main() {

	fmt.Println("iota and const of four years: ", a, b, c, d)

}
