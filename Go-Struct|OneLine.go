package main

import (
	"fmt"
)

func main() {
	aCar := struct{ name string }{name: "Nissan GTR R34"}
	fmt.Println(aCar)

}
