/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/4 3:18 下午
* @Desc: a
 */

package main

import (
	"bytes"
	"fmt"
	"testing"
)

//func TestAppend(t *testing.T) {
//	i :=1
//	i++
//	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
//	fmt.Println(strconv.Itoa(100))
//	fmt.Println(strconv.FormatInt(100, 10))
//	c := [3]int{1, 2, 3}
//	cc := c[:]
//	cc = append(cc, 1)
//	a := []int{1, 2, 3,}
//	b := a[2:]
//	fmt.Println(a)
//	fmt.Println(b)
//	a = append(a, 60)
//	b[0] = 10
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(cap(a))
//	fmt.Println(cap(b))
//
//	b = append(b, 20)
//	b[0] = 30
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(cap(a))
//	fmt.Println(cap(b))
//}
//
//func TestDilatation(t *testing.T) {
//	a := make([]int, 1023)
//	fmt.Println("-- less then 1023 --")
//	fmt.Printf("before cap : %d \n", cap(a))
//	a = append(a, 1)
//	fmt.Printf("after cap : %d \n", cap(a))
//	b := make([]int, 1024)
//	fmt.Println("-- more then 1024 --")
//	fmt.Printf("before cap : %d \n", cap(b))
//	b = append(b, 1)
//	fmt.Printf("after cap : %d \n", cap(b))
//}
//
//func TestCreateSlice(t *testing.T) {
//	var a []int
//	b := []int{}
//	c := make([]int, 0)
//	fmt.Println(cap(a))
//	fmt.Println(cap(b))
//	fmt.Println(cap(c))
//}
//
//func TestCopy(t *testing.T) {
//	a := []int{1,2,3}
//	b := make([]int, 4)
//	copy(b, a)
//	a[0] = 2
//	fmt.Println(a)
//	fmt.Println(b)
//}
var src = bytes.Repeat([]byte("golang"), 1024*10)
func BenchmarkAppend(b *testing.B) {
	for i:=0; i < b.N; i++ {

		_ = append([]byte(nil), src...)
	}
}

	func BenchmarkMakeCopy(b *testing.B) {

	for i:=0; i < b.N; i++ {
		aa := make([]byte, len(src))
		copy(aa, src)

	}
}

func TestWRPanic(t *testing.T) {
	s := []int{1,2,3}
	for i :=0;  i < 1000; i++ {
		go func(t int) {
			if t % 10 == 0 {
				s[0] = 2
			} else {
				fmt.Println(s[0])
			}
		}(i)
	}
}

type Test struct {
	Name string
}

func TestList(t *testing.T) {

	list := []*Test{&Test{Name:"22222"},&Test{Name:"3333"}}
	list2 := []string{}
	for _, test := range list {
		fmt.Println(&test)
		fmt.Println(test.Name)
		list2 = append(list2, test.Name)
		test.Name = test.Name + "asdfaaf"

	}
	for _, a := range list {
		fmt.Println(a.Name)
	}
	fmt.Println(list)
	fmt.Println(list2)

}