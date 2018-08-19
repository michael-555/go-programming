package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 30
		c <- 31
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
}
