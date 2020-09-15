/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/27 11:14 上午
* @Desc: a
 */

package channel

import (
	"fmt"
	"sync"
)

func IsClosed(ch chan int) bool {
	select {
	case v := <-ch:
		fmt.Println(v)
		return true
	default:
	}
	return false
}

func SafeClosed(ch chan int) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()
	close(ch)
	return true
}

type MyChanOnce struct {
	C chan int
	closed bool
	once sync.Once
}

func NewChan() *MyChanOnce {
	return &MyChanOnce{
		C: make(chan int),
	}
}

func(mc *MyChanOnce) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

type MyChanMutex struct {
	C chan int
	closed bool
	mutex sync.Mutex
}
func NewChanMutext() *MyChanMutex {
	return &MyChanMutex{
		C: make(chan int),
	}
}

func(mc *MyChanMutex) SafeClose() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
}

func(mc *MyChanMutex) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}
