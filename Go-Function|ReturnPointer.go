package main

import (
	"fmt"
)

func main() {
	s := sum(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum is", *s)
}
func sum(values ...int) *int {
	fmt.Println(values)
	sum := 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return &sum
}
