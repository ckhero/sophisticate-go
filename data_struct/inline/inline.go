package main

/**
内联
go build -gcflags=-m=2 inline.go
 */

func main()  {

}
//go:noinline
func Small(a, b int) int{
	return a + b
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
