/**
* @Author: 人从众[ckhero]
* @Date: 2020/8/26 6:56 下午
* @Desc: a
 */

package go_way

type Ints []int

func(i Ints) Iterator() <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range i {
			c <- v
		}
		close(c)
	}()
	return c
}

