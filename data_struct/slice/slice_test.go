/**
* @Author: 人从众[ckhero]
* @Date: 2020/9/4 3:18 下午
* @Desc: a
 */

package slice

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	a := []int{1, 2, 3}
	b := a[2:]
	fmt.Println(a)
	fmt.Println(b)
	a = append(a, 60)
	b[0] = 10
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(cap(a))
	fmt.Println(cap(b))

	b = append(b, 20)
	b[0] = 30
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(cap(a))
	fmt.Println(cap(b))

}

func TestDilatation(t *testing.T) {
	a := make([]int, 1023)
	fmt.Println("-- less then 1023 --")
	fmt.Printf("before cap : %d \n", cap(a))
	a = append(a, 1)
	fmt.Printf("after cap : %d \n", cap(a))
	b := make([]int, 1024)
	fmt.Println("-- more then 1024 --")
	fmt.Printf("before cap : %d \n", cap(b))
	b = append(b, 1)
	fmt.Printf("after cap : %d \n", cap(b))
}

func TestCreateSlice(t *testing.T) {
	var a []int
	b := []int{}
	c := make([]int, 0)
	fmt.Println(cap(a))
	fmt.Println(cap(b))
	fmt.Println(cap(c))
}