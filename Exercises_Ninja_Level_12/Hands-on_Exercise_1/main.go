package main

import (
	"fmt"

	"github.com/MikaelOttosson/Golang/GreaterCommons_golang/Exercises_Ninja_Level_12/Hands-on_Exercise_1/dog"
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
}
