package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

func openFile()  {
	//_, err := ioutil.ReadFile("/Users/ckhero/php-5.6.40.tar.gz") // just pass the file name
	_, err := ioutil.ReadFile("/Users/ckhero/sophisticate/sophiticate/IPZ-933-C.mp4") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("OPEN FILE")

}
func main() {
	runtime.GOMAXPROCS(1)
	go openFile()
	for  {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("i get schedule")
	}
}
