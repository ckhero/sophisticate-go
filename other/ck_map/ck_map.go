package ck_map

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

var keys []string

func SortRange(data map[string]interface{})  {

	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v)
	}
}



func TestMap() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(200)
	m := make(map[int]int)
	for j := 0; j < 100; j++ {
		go func() {
			fmt.Println(j)
			if _, ok := m[j]; ok {
				fmt.Println("get data", j)
			} else {
				fmt.Println("no data")
			}
			wg.Done()
		}()
	}

	for j :=0; j < 100; j ++ {
		go func(t int) {
			m[t] = t
		}(j)
		wg.Done()

	}
	wg.Wait()
}

