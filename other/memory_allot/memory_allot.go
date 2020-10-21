package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var name  int64= 18801613198
	fmt.Println(name)
	fmt.Println(1 & 0x07)
	fmt.Println(1 << (1 & 0x07))
	fmt.Println(100 & 7)
	//Demo()
}

type A interface {
	Test() int
}
func Demo() int {
	c := make(chan *int)
	bb := 22

	c<-&bb
	//s := []byte("")
	//
	//s1 := append(s, 'a')
	//s2 := append(s, 'b')
	//
	//// 如果有此行，打印的结果是 a b，否则打印的结果是b b
	//// fmt.Println(s1, "===", s2)
	//fmt.Println(string(s1),)
	//fmt.Println(string(s2))
	a := (*Test)(test())
	One()
	fmt.Println(a.Name)
	return 	2

}
// 避免逃逸
func test() unsafe.Pointer {
	return noescape(unsafe.Pointer(&Test{
		Name: "name",
	}))
}

func One() {
	one := 1
	two := &one
	fmt.Println(two)
}
//
//func Two() {
//	two := make(chan *int)
//	twoVal := 1
//	two <- &twoVal
//}
////go:noinline
//func Three() {
//	l := 10
//	three := make([]unsafe.Pointer, l)
//
//	threecVal := 1
//	three = append(three ,  noescape(unsafe.Pointer(&threecVal)))
//}
type Test struct {
	Name string
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}


//type Stringer interface {
//	String() string
//}
//
//func Four() string{
//	var any interface{}
//	if v, ok := any.(Stringer); ok {
//		return v.String()
//	}
//	return ""
//}

