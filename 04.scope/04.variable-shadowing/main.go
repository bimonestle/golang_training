package main

import (
	"fmt"
)

func max(x int) int {
	return 42 + x
}

func main() {
	// max := max(7)
	plus := max(1)
	// fmt.Println(max) // max is now the result, not the fnction
	fmt.Println(plus)
	fmt.Println(max(8))
}

// don't do this; bad coding practice to shadow variables
