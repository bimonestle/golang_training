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
	fmt.Println(*b) // 43

	*b = 19        // b says, "The value at this address, changed to 42"
	fmt.Println(a) // 19

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at the memory address
	// this makes our program more performant
}
