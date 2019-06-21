package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// no acess to x
	// this doesn't compile
	fmt.Println(x)
}
