package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("Deferred Statement") //Defer line runs before panic statement
	panic("Error")
	fmt.Println("End")

}
