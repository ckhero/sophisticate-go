/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/27 3:40 下午
* @Desc: a
 */

package channel

import (
	"fmt"
	"sync"
)

func OneToN() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	ch := make(chan int, 100)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	} ()
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for v := range ch {
				fmt.Printf("name: %d,no: %d \n", i, v)
			}
		}()
	}

	wg.Wait()
}