package main

import "fmt"

func main() {

	s := "Hi, Google!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) // String converted into slice of bytes UTF8
	fmt.Println(bs)
	fmt.Printf("%T\n", bs) // defaults to type uint8

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n ", i, v)
	}

}
