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

	pop, boolean := statePopulations["Texas"] // Used since map outputs unknown keys with value 0, Here Boolean holds either True or False, If True, key exists in map, else it doesn't.
	fmt.Println(pop, boolean)
}
