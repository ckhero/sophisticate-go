package main

import (
	"fmt"
	"runtime"
	"time"
)

func deadLoop() {
	for {
		add(2, 3)

	}
}

func add(a, b int)  int {
	fmt.Println("I get scheduled two!")

	return a + b
}

func main() {
	runtime.GOMAXPROCS(1)
	go deadLoop()
	for {
		time.Sleep(time.Nanosecond)
		add(2, 3)

		fmt.Println("I get scheduled one!")
	}
}