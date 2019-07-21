package main

import (
	"fmt"
)

func main() {
	switch "Medhi" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your name starts with M")
	case "Julian", "Sushanti":
		fmt.Println("Wassup Julian / Sushanti")
	}
}
