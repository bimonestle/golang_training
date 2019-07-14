package main

import (
	"fmt"
)

func zero(z *int) {
	fmt.Println(z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println(x)  // x is 5
	fmt.Println(&x) // memory address of 5
	zero(&x)        // memory address of 5
	fmt.Println(x)  // x is 0
}
