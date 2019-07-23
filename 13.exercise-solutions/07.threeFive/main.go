package main

import (
	"fmt"
)

func main() {
	var totNumb int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println(i)
			totNumb = totNumb + i
		}
	}
	fmt.Println(totNumb)

}
