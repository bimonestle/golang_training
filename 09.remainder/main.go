package main

import (
	"fmt"
)

func main() {
	x := 12 % 3
	fmt.Println(x)

	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println(i, "Odd")
		} else {
			fmt.Println(i, "Even")
		}
	}
}
