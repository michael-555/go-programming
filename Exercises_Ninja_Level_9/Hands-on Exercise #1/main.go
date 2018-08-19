package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Begin cpu", runtime.NumCPU())
	fmt.Println("Begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Thing 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Thing 2")
		wg.Done()
	}()

	fmt.Println("Mid cpu", runtime.NumCPU())
	fmt.Println("Mid gs", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("about to exiting...")

	fmt.Println("End cpu", runtime.NumCPU())
	fmt.Println("End gs", runtime.NumGoroutine())
}
