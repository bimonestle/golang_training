package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	BIMO := "BIMO"
	fmt.Println(BIMO)
	fmt.Printf("%T \n", BIMO)
	fmt.Println([]byte(string(BIMO))) // rune

	TEGUH := "Teguh"
	fmt.Println([]byte(string(TEGUH))) // rune
}
