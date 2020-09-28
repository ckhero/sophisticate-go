package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			go func() {
				for {
					fmt.Println(1)
				}
			}()
		}
	}()
	http.ListenAndServe("0.0.0.0:6060", nil)
}