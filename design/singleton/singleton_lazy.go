package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct {}
var ins *Singleton
var mu sync.Mutex
var once sync.Once

/**
懒汉式创建，需要使用的时候创建
 */
func GetIns() *Singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			fmt.Println("create singleton")
			ins = &Singleton{}
		}
	}
	return ins
}

func GetInsOnce() *Singleton {
	once.Do(func() {
		ins = &Singleton{}
	})
	return ins
}