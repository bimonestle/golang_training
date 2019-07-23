package main

import (
	"fmt"
)

func main() {
	var smNumber int
	var lgNumber int
	fmt.Print("Please enter small number: ")
	fmt.Scan(&smNumber)
	fmt.Print("Please enter large number: ")
	fmt.Scan(&lgNumber)
	fmt.Println("The remainder of", lgNumber, "%", smNumber, "is ", lgNumber%smNumber)
}
