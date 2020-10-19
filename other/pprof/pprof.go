package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var datas []string
var datas2 []string

func Add(data string) string {
	//time.Sleep(time.Nanosecond * 100)
	sdata := []byte(data)
	str := string(sdata)
	datas = append(datas, str)
	return str
}

func Add2(data string) string {
	//time.Sleep(time.Nanosecond * 200)

	sdata := []byte(data)
	str := string(sdata)
	datas2 = append(datas2, str)
	return str
}

var m1 sync.Mutex
func main() {
	runtime.GOMAXPROCS(2)
	go func() {
		fmt.Println("1 start")

		m1.Lock()
		time.Sleep(time.Second * 10)
		fmt.Println("1 end")

	}()
	go func() {

		m1.Lock()
		fmt.Println("2 start")

	}()
	go func() {

		for {

			Add("https://segmentfault.com/a/1190000016412013")
			Add2("2")
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}