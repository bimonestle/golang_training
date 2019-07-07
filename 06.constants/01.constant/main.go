package main

import (
	"fmt"
)

const p string = "death & taxes"
const r = "Constant r"

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println(p, r, q)
}

// a CONSTANT is a simple unchanging value
