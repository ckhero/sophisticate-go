package main

var test int

const (
	mutexLocked =  1<<iota
	mutexLocked2
)


func main() {
	mm  := make(map[string]string,111111111111121)
	mm["2"]="2"
	//fmt.Println(mm)
	//for i:= 0; i <  100000; i++ {
	//	mm[i]=i
	//}
	//test1()
	//test2()
	//test3(&mm)
}

//func test1() *int{
//	aa := 0
//	return &aa
//}
//func test2() int{
//	bb := 0
//	return bb
//}

//func test3(c *map[string]int) {
//	fmt.Println(c)
//}