package number__

import "testing"

func TestOne(t *testing.T) {
	One()
	Two()
	Three()

	test1()
	test2()
}


func BenchmarkTwo(b *testing.B) {
	Two()
}

func BenchmarkThree(b *testing.B) {
	Three()
}