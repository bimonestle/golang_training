package main

import (
	"fmt"
)

func main() {
	switch "Daniel" {
	// switch "Mhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?") // will print this if the switch statement doesn't match with any
	}
}
