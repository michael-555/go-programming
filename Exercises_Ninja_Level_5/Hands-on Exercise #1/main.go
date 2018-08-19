package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	icecream  []string
}

func main() {

	p1 := person{
		firstname: "James",
		lastname:  "Bond",
		icecream: []string{
			"Vanilla",
			"Martini",
			"Rum",
		},
	}
	p2 := person{
		firstname: "Miss",
		lastname:  "Moneypenny",
		icecream: []string{
			"Strawberry",
			"Banana",
			"capuccino",
		},
	}

	// Mikael extra
	p3 := person{"Ian", "Flemming", []string{"Whiskey"}}
	fmt.Println("Name: ", p3.firstname, p3.lastname)
	fmt.Print("Favorite flavors: ")
	for _, v := range p3.icecream {
		fmt.Printf("%v ", v)
	}

	fmt.Printf("\n\nName: %v %v\n", p1.firstname, p1.lastname)
	fmt.Print("Favorite flavors: ")
	for _, v := range p1.icecream {
		fmt.Printf("%v ", v)
	}

	fmt.Printf("\n\nName: %v %v\n", p2.firstname, p2.lastname)
	fmt.Print("Favorite flavors: ")
	for _, v := range p2.icecream {
		fmt.Printf("%v ", v)
	}

}
