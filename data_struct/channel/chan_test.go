package channel

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

func TestIsClosed(t *testing.T) {
	ch := make(chan int)
	fmt.Println(IsClosed(ch))
	close(ch)
	fmt.Println(IsClosed(ch))
}

func TestSafeClosed(t *testing.T) {
	ch := make(chan int)
	fmt.Println(SafeClosed(ch))
	fmt.Println(SafeClosed(ch))
}

func TestMyChanOnce_SafeClose(t *testing.T) {
	ch := NewChan()
	ch.SafeClose()
	ch.SafeClose()
}

func TestMyChanMutex_SafeClose(t *testing.T) {
	ch := NewChanMutext()
	fmt.Println(ch.IsClosed())
	ch.SafeClose()
	fmt.Println(ch.IsClosed())
}