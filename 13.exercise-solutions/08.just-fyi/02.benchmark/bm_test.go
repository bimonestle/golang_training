package main

import (
	"fmt"
	"testing" // https://golang.org/pkg/testing
)

func BenchmarkHello(b *testing.B) {
	counter := 0
	for i := 0; i < b.N; i++ {
		// if i%3 == 0 || i%5 == 0 {
		// 	counter += i
		// }
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}
