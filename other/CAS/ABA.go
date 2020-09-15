/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/18 4:12 下午
* @Desc: ABA问题，线程一修改 v 1 -> 2 -> 1,线程2 修改v 1->3,实际上线程2修改时候的1已经不是原来的1了
 */

package CAS

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func ABA() {
	var num int32 = 20
	var version int32 = 0
	a := version
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		atomic.CompareAndSwapInt32(&num, 20, 30)
		atomic.AddInt32(&version, 1)
		fmt.Printf("num1:%d \n", atomic.LoadInt32(&num))
		atomic.CompareAndSwapInt32(&num, 30, 20)
		atomic.AddInt32(&version, 1)
		fmt.Printf("num2:%d \n", atomic.LoadInt32(&num))
		wg.Done()
	} ()

	go func() {
		time.Sleep(time.Second)
		if atomic.CompareAndSwapInt32(&version, a, a + 1) {
			if atomic.CompareAndSwapInt32(&num, 20, 40) {
				fmt.Printf("num3:%d \n", atomic.LoadInt32(&num))
			} else {
				fmt.Printf("fail - num3:%d \n", atomic.LoadInt32(&num))
			}
			fmt.Printf("version:%d \n", version)
		} else {
			fmt.Printf("version:%d \n", version)
		}
		wg.Done()
	} ()
	wg.Wait()
}
