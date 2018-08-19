package main

import (
	"fmt"
)

func main() {

	emp1 := struct {
		name, department string
		salery, age      int
		friends          map[string]int
		favDrinks        []string
	}{
		name:       "Mike",
		department: "Finance",
		salery:     500,
		age:        30,
		friends: map[string]int{
			"Jenny": 333,
			"Emmie": 222,
		},
		favDrinks: []string{
			"Rum",
			"Coke",
			"Water",
		},
	}

	fmt.Printf("Name: %v\n", emp1.name)
	fmt.Printf("Department: %v\n", emp1.department)
	fmt.Printf("Salery: %v\n", emp1.salery)
	fmt.Printf("Age: %v\n", emp1.age)
	fmt.Print("Friends: ")
	for k, v := range emp1.friends {
		fmt.Printf("%v, Phonenumber: %v ", k, v)
	}
	fmt.Print("\nFavorite drinks: ")
	for _, val := range emp1.favDrinks {
		fmt.Printf("%v ", val)
	}
}
