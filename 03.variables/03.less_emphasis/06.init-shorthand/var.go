package main

import (
	"fmt"
	"math"
)

func main() {
	// you can only do this inside a func
	message := "Hello Init Shorthand"
	a, b, c := 1, true, "b"
	d := 4.05
	e := false
	f := "Bimo"
	g := math.Pow(2, 8)

	fmt.Println(message, a, b, c, d, e, f, g)
}
