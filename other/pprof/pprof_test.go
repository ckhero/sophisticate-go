package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i :=0; i < b.N; i ++ {
		Add("https://segmentfault.com/a/1190000016412013")
		Add2("23")
	}
}