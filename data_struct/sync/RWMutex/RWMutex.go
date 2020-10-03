package main

import (
	"fmt"
	"sync"
)

func main()  {
	rw := sync.RWMutex{}
	rw.RLock()
	rw.RLock()
	rw.RUnlock()
	rw.Lock()
	fmt.Println(2222)
}
