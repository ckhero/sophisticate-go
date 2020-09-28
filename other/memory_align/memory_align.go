/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/26 3:57 下午
* @Desc: 内存对齐
 */

package memory_align

import (
	"fmt"
	"unsafe"
)

type SizeOfF struct {
	A byte  // 1
	B int32 // 8
	C byte  // 1]
	D int32 // 8
}

type SizeOfF2 struct {
	A byte  // 1
	C byte  // 1
	B int64 // 8
	D string // 8
}

func OptimizeMemory() {
	fmt.Println("-----SizeOfF------")
	fmt.Println("")
	fmt.Printf("A size: %d,align: %d, offset: %d \r\n", unsafe.Sizeof(SizeOfF{}.A), unsafe.Alignof(SizeOfF{}.A), unsafe.Offsetof(SizeOfF{}.A))
	fmt.Printf("B size: %d,align: %d, offset: %d \r\n", unsafe.Sizeof(SizeOfF{}.B), unsafe.Alignof(SizeOfF{}.B), unsafe.Offsetof(SizeOfF{}.B))
	fmt.Printf("C size: %d,align: %d, offset: %d \r\n", unsafe.Sizeof(SizeOfF{}.C), unsafe.Alignof(SizeOfF{}.C), unsafe.Offsetof(SizeOfF{}.C))
	fmt.Printf("D size: %d,align: %d, offset: %d \r\n", unsafe.Sizeof(SizeOfF{}.D), unsafe.Alignof(SizeOfF{}.D), unsafe.Offsetof(SizeOfF{}.D))
	fmt.Println("")
	fmt.Println("-----SizeOfF2------")
	fmt.Println("")
	fmt.Printf("A size: %d,align: %d, offset: %d \r\n", unsafe.Sizeof(SizeOfF2{}.A), unsafe.Alignof(SizeOfF2{}.A), unsafe.Offsetof(SizeOfF2{}.A))
	fmt.Printf("B size: %d,align: %d, offset: %d \r\n", unsafe.Sizeof(SizeOfF2{}.B), unsafe.Alignof(SizeOfF2{}.B), unsafe.Offsetof(SizeOfF2{}.B))
	fmt.Printf("C size: %d,align: %d, offset: %d \r\n", unsafe.Sizeof(SizeOfF2{}.C), unsafe.Alignof(SizeOfF2{}.C), unsafe.Offsetof(SizeOfF2{}.C))
	fmt.Printf("D size: %d,align: %d, offset: %d \r\n", unsafe.Sizeof(SizeOfF2{}.D), unsafe.Alignof(SizeOfF2{}.D), unsafe.Offsetof(SizeOfF2{}.D))
	fmt.Printf("size of fdsaa %d \n", unsafe.Sizeof("fdsaa"))
	fmt.Printf("size of fdsaafdsaafdsaafdsaa %d \n", unsafe.Sizeof("fdsaafdsaafdsaafdsaa"))
	fmt.Printf("size of int %d \n", unsafe.Sizeof(1))
	fmt.Printf("size of int8 %d \n", unsafe.Sizeof(int8(1)))
	fmt.Printf("size of int16 %d \n", unsafe.Sizeof(int16(1)))
	fmt.Printf("size of int32 %d \n", unsafe.Sizeof(int32(1)))
	fmt.Printf("size of int64 %d \n", unsafe.Sizeof(int64(1)))
	fmt.Printf("size of true %d \n", unsafe.Sizeof(true))
	fmt.Printf("size of byte %d \n", unsafe.Sizeof(byte(1)))
	fmt.Printf("D size: %d,align: %d \r\n", unsafe.Sizeof(SizeOfF{}), unsafe.Alignof(SizeOfF{}))

}