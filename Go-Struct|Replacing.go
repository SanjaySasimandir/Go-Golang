package main

import (
	"fmt"
)

type Car struct {
	name       string
	year       int
	pastOwners []string
}

func main() {
	aCar := Car{
		name: "Nissan GTR R34",
		year: 1998,
		pastOwners: []string{
			"David",
			"Smith",
		},
	}

	fmt.Println(aCar)
	fmt.Println(aCar.name)
	aCar.name = "Toyota Supra MKIV"
	fmt.Println(aCar.name)
	fmt.Println(aCar)

}
