package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main() {
	sync.Mutex{}
	s := make([]int, 10000000)
	ss := new(map[string]string)
	s[1] = 2
	fmt.Println(*ss)
}


type Test struct {
}
// 避免逃逸
func test() unsafe.Pointer {
	return noescape(unsafe.Pointer(&Test{}))
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}
