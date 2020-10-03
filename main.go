package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"main/design/chain"
	"time"
)

func main()  {
	var g singleflight.Group
	for i :=0; i < 2; i ++  {
		go func() {
			v,err, s :=g.Do("test", func() (i interface{}, err error) {
				fmt.Println("dsf")
				return "aaaa", errors.New("error")
			})
			fmt.Println(v,err,s)
		}()
	}

	time.Sleep(time.Second)
	v1 := chain.NewTeacher()
	v2 := chain.NewHeadermaster()
	v1.SetNext(v2)
	v := v1
	v.Exec(1)
	v.Exec(3)
	v.Exec(2)
	v.Exec(4)
}