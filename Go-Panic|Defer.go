package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	defer panic("Error") //Panic Statements can be deferred
	fmt.Println("End")

}
