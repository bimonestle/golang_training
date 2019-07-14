package main

import (
	"fmt"
)

func main() {
	a := 43
	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x...

	var b *int = &a
	fmt.Println(b)  // 0x...
	fmt.Println(*b) // 43. This is dereferencing

	// b is a int pointer
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add an * in front of b
	// this is known as dereferencing
	// the * is an operator in this case
}
