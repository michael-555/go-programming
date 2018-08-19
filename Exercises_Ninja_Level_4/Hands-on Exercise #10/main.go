package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`Ian Flemming`] = []string{`steaks`, `cigars`, `espionage`} // adds a record to m map

	fmt.Println(m)

	delete(m, `Ian Flemming`) // deletes a map record

	for k, v := range m {
		fmt.Println("This is the record for:", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
