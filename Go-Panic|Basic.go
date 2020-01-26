package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	panic("Error")
	fmt.Println("End")

}
