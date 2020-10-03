package timeout

import "time"

/**
https://studygolang.com/articles/27402
 */
func one() {
	<-time.After(time.Second)
}

func two()  {
	timer := time.NewTimer(time.Second * 10)
	<- timer.C
}