package number__

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func One() {
	total, sum := 0, 0
	var wg sync.WaitGroup

	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		sum += i
		go func(tmp int) {
			defer wg.Done()
			total += tmp
		}(i)
	}
	wg.Wait()
	fmt.Printf("one: total:%d sum %d \n", total, sum)
}


func Two() {
	total, sum := 0, 0
	t := time.Now()
	var lock sync.Mutex
	var wg sync.WaitGroup
	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		sum += i

		go func(tmp int) {
			defer func() {
				lock.Unlock()
				wg.Done()
			}()
			lock.Lock()
			total += tmp
		}(i)
	}
	wg.Wait()

	fmt.Printf("two: total:%d sum %d \n", total, sum)
	fmt.Printf("two: cost %d \n", time.Now().Sub(t))

}

func Three() {
	t := time.Now()
	sum := 0
	var total uint64 = 0
	var wg sync.WaitGroup

	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		sum += i
		go func(tmp int) {
			defer wg.Done()
			atomic.AddUint64(&total, ^uint64(tmp - 1))
			//for {
			//	old := atomic.LoadInt64(&total)
			//	if atomic.CompareAndSwapInt64(&total, old, (old + int64(tmp))) {
			//		break
			//	}
			//}
		}(i)
	}
	wg.Wait()

	fmt.Printf("three: total:%d sum %d \n", total, sum)
	fmt.Printf("three: cost %d \n", time.Now().Sub(t))

}


func test1() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	count := int64(0)
	t := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			count++
			wg.Done()
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Printf("test1 花费时间：%d, count的值为：%d \n", time.Now().Sub(t), count)
}

// sync.atomic
func test2() {
	var wg sync.WaitGroup
	count := int64(0)
	t := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("test2 花费时间：%d, count的值为：%d \n", time.Now().Sub(t), count)
}