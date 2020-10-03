package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex
var c = sync.NewCond(&m)
func main() {
	DemoTwo()
	time.Sleep(time.Second)
}
func DemoTwo() {
	var m int
	go func() {
		for m != 1 {
			c.L.Lock()
			fmt.Println("one wait")
			c.Wait()
			fmt.Println(m)
			c.L.Unlock()
		}
		fmt.Println("one quit")
	}()
	go func() {
		for m != 4 {
			c.L.Lock()
			fmt.Println("two wait")
			c.Wait()
			fmt.Println(m)

			c.L.Unlock()
		}
		fmt.Println("two quit")
	}()
	for m = 0; m < 5; m++ {
		time.Sleep(time.Second)
		c.Broadcast()
	}
}


func DemoOne() {
	running := make(chan bool, 5)
	for i := 0; i < 5; i ++ {
		go syncCondTest(i, running)
	}
	for i :=0 ;i < 5; i ++ {
		<-running
	}
	fmt.Println("all goroutine runnningn")
	c.Signal()
	//c.Broadcast()
	time.Sleep(time.Second)
}
func syncCondTest(i int, running chan bool) {
	fmt.Println("in - go : %d", i)

	c.L.Lock()
	defer func() {
		c.L.Unlock()
	}()
	fmt.Println("wait - go : %d", i)
	running <- true
	c.Wait()
	fmt.Println("quit - go : %d", i)
}


