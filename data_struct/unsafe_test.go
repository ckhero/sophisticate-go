package data_struct

import (
	"fmt"
	"testing"
	"unsafe"
)

type Test1 struct {
	Name string
}

type Test12 struct {
	Name string
}
func TestUnsafe(t *testing.T) {
	//a := 1
	//aptr := &a
	//fmt.Printf("size of aptr is %d \n", unsafe.Sizeof(aptr))

	a := &Test1{Name:"aaaa"}
	b := (*Test12)(unsafe.Pointer(a))
	fmt.Println(b)
}