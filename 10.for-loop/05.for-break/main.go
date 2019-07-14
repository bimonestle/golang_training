package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}

	for i := 11; i < 16; i++ {
		fmt.Println(i)
		if i == 14 {
			break
		}
	}
}
