package limit_ip

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func TestLimitIp_IsInvalid(t *testing.T) {
	limitIp := &LimitIp{}
	var res int64 = 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i ++ {

		if i == 3 {
			//time.Sleep(time.Second * 11)
		}
		//time.Sleep(time.Second)
		for j := 0; j < 100; j ++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				if  limitIp.IsInvalid(strconv.Itoa(j)) {
					atomic.AddInt64(&res, 1)
				}
			}()
		}

	}
	wg.Wait()
	fmt.Printf("res=%d",res)

}