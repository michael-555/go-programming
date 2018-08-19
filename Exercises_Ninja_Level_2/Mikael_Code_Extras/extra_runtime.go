package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)   // Print system env
	fmt.Println(runtime.GOARCH) // Print system arch.
}
