package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 340; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
	fmt.Printf("%v \n", "BIMO")
}
