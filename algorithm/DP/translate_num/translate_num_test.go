package main

import "testing"

func BenchmarkTranslateNum(b *testing.B) {
	for i :=0; i < b.N; i ++ {
		TranslateNum(1068385902)
		TranslateNum2(1068385902)
	}
}