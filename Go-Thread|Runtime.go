package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Threads:", runtime.GOMAXPROCS(-1))
}
