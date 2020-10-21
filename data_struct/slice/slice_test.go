/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/4 3:18 下午
* @Desc: a
 */

package slice

import (
	"bytes"
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