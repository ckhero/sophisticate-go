package main

import "fmt"

type Test struct {
	Name string
}

type Foo struct {
	Bar string
}
/**
https://studygolang.com/articles/25465?fr=sidebars
 */
func main() {
	arr := []int{1, 2, 3}

	myMap := make(map[int]*int)
	for i, v := range arr {
		vTmp := v
		myMap[i] = &vTmp
	}

	for _, v := range myMap {
		fmt.Println(*v)
	}
}
