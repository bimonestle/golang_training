package main

import (
	"fmt"
	"testing" // https://golang.org/pkg/testing
)

func BenchMarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello")
	}
}

// run this at command:
// go test -bench='.*'
