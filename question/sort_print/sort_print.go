package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main()  {
	//DoPrint()
	UseChannel()
}

//func UseCond() {
//	s := -1
//	var l sync.Mutex
//	c := sync.NewCond(&l)
//	m, n := 0, 1
//	go func() {
//		for s == m  {
//			fmt.Println(m)
//			m += 2
//		}
//	}()
//}
func UseChannel() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go func() {
		for i := 0; i < 10; i += 2 {
			<- ch1
			fmt.Println(i)
			ch2 <- true
		}
	}()

	go func() {
		for i := 1; i < 10; i += 2 {
			<- ch2
			fmt.Println(i)
			ch1 <- true
		}
	}()
	ch1 <- true
	time.Sleep(time.Second)
}

func SpinkLock(i int64, f func(), count *int64) {
	for  {
		if i == atomic.LoadInt64(count) {
			f()
			atomic.AddInt64(count, 1)
			break
		}
		time.Sleep(time.Nanosecond)
	}
}

func DoPrint() {
	var count int64 = 0
	for i :=0; i < 10; i ++ {
		f := func() {
			fmt.Println(atomic.LoadInt64(&count))
		}
		go SpinkLock(int64(i), f, &count)
	}
	time.Sleep(time.Second)
}
