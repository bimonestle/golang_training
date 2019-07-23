package main

import (
	"fmt"
)

func main() {
	// we can initializea variable in the conditional statements; food
	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	c := "at"
	if len(c) > 3 {
		fmt.Println(b)
	} else {
		fmt.Println(!b)
	}
}
