package main

import (
	"fmt"
)

func main() {
	b := true
	if food := "Chocolate"; !b {
		fmt.Println("?")
	} else {
		fmt.Println(food)
	}
	// fmt.Println(food)  // the scope of variable food encloses the conditional statements
}
