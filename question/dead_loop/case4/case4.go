package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func deadLoop() {
	for {
		go func() {
			dummy()
			fmt.Println(			runtime.NumGoroutine())
		}()
	}
}

func dummy() {
	add(2, 3)
}
//go:noinline
func add(a, b int)  int {
	return a + b

}

func main() {
	go deadLoop()
	//for {
	//	time.Sleep(time.Second)
	//	add(3, 4)
	//}
	http.ListenAndServe("0.0.0.0:6060", nil)
}