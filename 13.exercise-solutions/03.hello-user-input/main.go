package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Hello! Please enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello!", name)
}
