package t_test

import "testing"

/**
内联
go build -gcflags=-m=2 inline_test.go
 */

func main()  {

}

func Small(a, b int) int{
	return a + b
}
func BenchmarkSmall(b testing.B) {
	Small()
}
func large() string {
	a := "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	a += "a"
	return a
}
