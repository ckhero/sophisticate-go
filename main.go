/**
 *@Description
 *@ClassName log
 *@Date 2020/11/2 11:02 下午
 *@Author ckhero
 */

package main

import "fmt"

type Test struct {

}

func(t *Test) check() {
	fmt.Println(1111111)
}


type Test2 struct {
	Test
}

func(t *Test2) check() {
	t.Test.check()
	fmt.Println(2222)
}

func main() {
	t := &Test2{Test{}}
	t.check()
}