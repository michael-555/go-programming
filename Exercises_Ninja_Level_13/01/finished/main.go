package main

import (
	"fmt"

	"github.com/MikaelOttosson/Golang/GreaterCommons_golang/Exercises_Ninja_Level_13/01/starting-code/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
