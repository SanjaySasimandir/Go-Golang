package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California": 19516,
		"Illinois":   56516,
		"Texas":      556555,
		"Georgia":    5445454,
	}
	fmt.Println(len(statePopulations))
}
