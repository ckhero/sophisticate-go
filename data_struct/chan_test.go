package data_struct

import (
	"fmt"
	"sync"
	"testing"
)

func TestChan(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	fmt.Printf("cap:%d \n", cap(ch))
	fmt.Printf("size:%d \n", len(ch))
}

var ch = make(chan bool, 1)
func UserChan() {
	ch <- true
	<-ch
}
var mutex = sync.Mutex{}

func UserMutex() {
	mutex.Lock()
	mutex.Unlock()
}

func BenchmarkUseMutext(b *testing.B)  {
	for i := 0; i < 10000; i++ {
		UserMutex()
	}
}

func BenchmarkUseChan(b *testing.B)  {
	for i := 0; i < 10000; i++ {
		UserChan()
	}
}