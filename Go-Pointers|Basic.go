package main

import (
	"fmt"
)

func main() {
	var a int = 28
	var b *int = &a
	fmt.Println(a, *b)
	fmt.Println(a, b) //b=address of a
	a = 76
	fmt.Println(a, *b) //*b=value of a
}
