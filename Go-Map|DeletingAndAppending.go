package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California": 19516,
		"Illinois":   56516,
		"Texas":      556555,
	}
	statePopulations["Georgia"] = 5445454
	delete(statePopulations, "Texas")
	fmt.Println(statePopulations["California"])
	fmt.Println(statePopulations)
}
