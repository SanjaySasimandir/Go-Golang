package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("1") //Defer runs lines last, Priority = last>first
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	defer fmt.Println("5")

}
