package main

import (
	"fmt"
)

func main() {
	a := 43
	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x...

	var b = &a
	fmt.Println(b) // 0x...

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of the type *int
}
