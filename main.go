package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)



func main()  {
	m := sync.Map{}
	m.Store("key", "val")
	for i :=0; i < 100;  i ++ {
		go func() {
			fmt.Println(m.Load("key"))
		}()
		if i == 30 {
			m.Store("key", "val2")
		}
	}
	time.Sleep(time.Second)
}

func GetData(key string, val int, g *singleflight.Group) interface{} {
	v, _, _ :=g.Do(key, func() (i interface{}, err error) {
		// 处理逻辑
		fmt.Println("first", val)
		return val, nil
	})
	return v
}
