package main

import (
	"fmt"
	"math"

	"github.com/bimonestle/golang_training/03.variables/03.less_emphasis/07.All-together/mathfunc"
)

var a = "Stored in the variable a"                 // package scope
var b, c string = "Stored in var b", "Stored in c" // package scope
var d string                                       // package scope

func main() {
	d = "stored in D" // declaration above; assignment here; package scope
	var e = 42        // function scope - subsequent variables have func sc$
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "-n-"                           // double quotes
	o := `ooo`                           // back ticks
	p := 3.232
	// q := 42.38

	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h, i)
	fmt.Println(j, k, l, m)
	fmt.Println(n, o)
	fmt.Println(math.Floor(p))
	fmt.Println(math.Ceil(p))
	fmt.Println(math.Round(p))
	fmt.Println("P variable = ", mathfunc.Round(p, 0.5))
}
