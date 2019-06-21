package main

import (
	"fmt"
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 9 // error because x is undefined when invoked
}

var y = 12
