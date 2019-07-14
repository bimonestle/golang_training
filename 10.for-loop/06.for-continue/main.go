package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
	// It will print odd numbers only

	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
	// Doing the same thing
}
