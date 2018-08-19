package main

import (
	"fmt"

	"github.com/MikaelOttosson/Golang/GreaterCommons_golang/Exercises_Ninja_Level_13/02/02-code-finished/quote"
	"github.com/MikaelOttosson/Golang/GreaterCommons_golang/Exercises_Ninja_Level_13/02/02-code-finished/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
