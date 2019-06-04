package main

import (
	"fmt"
)

func main() {
	var message = "Hello Infer Mixed Up Type!"
	var a, b, c, d, e, f, g = 1, false, 3, 4.50, "35", "a", true

	fmt.Println(message, a, b, c, d, e, f, g)
	fmt.Println(message, "\n", a, "\n", b, "\n", c, "\n", d, "\n", e, "\n", f, "\n", g)
}
