package main

import (
	"fmt"

	"github.com/bimonestle/golang_training/04.scope/01.package-scope/02.visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
