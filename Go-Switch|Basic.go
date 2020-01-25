package main

import (
	"fmt"
)

func main() {
	switch 2 {
	case 1,4,5:
		fmt.Println("One")
	case 2,6,7:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}
}
