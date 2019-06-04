package main

import (
	"fmt"
)

func main() {
	var message = "Hello Infer Type!"
	var a, b, c = 1, 2, 3

	fmt.Println(message, a, b, c)
	fmt.Print(message)
	fmt.Print(a, b, c)
	fmt.Print(a, b, c, "\n")
	fmt.Print(message)
}
