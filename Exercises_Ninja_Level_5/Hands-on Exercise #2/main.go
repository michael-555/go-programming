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

	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	//fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.firstname, v.lastname)
		for i, val := range v.icecream {
			fmt.Println(i, val)
		}
		fmt.Println("--------------------------------")
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

	// Mikael extra
	p3 := person{"Ian", "Flemming", []string{"Whiskey"}}
	fmt.Println("Name: ", p3.firstname, p3.lastname)
	fmt.Print("Favorite flavors: ")
	for _, v := range p3.icecream {
		fmt.Printf("%v ", v)
	}

}
