package main

import (
	"fmt"
)

func main() {
	greet("Jane ", "Doe")
	greet2(24, "Mary")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}

func greet2(age int, name string) {
	fmt.Println(age, name)
}
