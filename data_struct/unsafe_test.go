package data_struct

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	a := 1
	aptr := &a
	fmt.Printf("size of aptr is %d \n", unsafe.Sizeof(aptr))
}