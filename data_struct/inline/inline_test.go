package main

import "testing"

func BenchmarkSmall(b *testing.B) {
	for i :=0; i < 100000; i++ {
		Small(1, 2)
	}
}
