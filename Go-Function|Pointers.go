package main

import (
	"fmt"
)

func main() {
	greeting := "Hello"
	name := "Everyone"
	printMessage(&greeting, &name)
	fmt.Println(name)
}
func printMessage(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Bob"
	fmt.Println(*name)
}
