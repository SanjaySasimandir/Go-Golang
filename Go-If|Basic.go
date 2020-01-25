package main

import (
	"fmt"
)

func main() {
	nissan := 50
	supra := 42
	if nissan == supra {
		fmt.Println("=")
	}
	if nissan > supra {
		fmt.Println(">")
	}

	if nissan < supra {
		fmt.Println("<")
	}
	fmt.Println(nissan == supra, nissan > supra, nissan < supra, nissan&^supra, nissan&supra, nissan|supra)
}
