package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println(5000000 * 8)
	ctx1, _ := context.WithCancel(context.Background())
	ctx, _ := context.WithTimeout(ctx1, time.Second)
	go Test(ctx)
	//cancel()
	go func() {
		for {
		select {
		case <-ctx.Done():
			fmt.Println("end")
			return
		}
	}}()
	time.Sleep(time.Second * 22)

}

func Test(ctx context.Context) {
	defer func() {
		fmt.Println("test exit")
	}()
	for  {
		select {
			case <-ctx.Done():
				return
		default:
			time.Sleep(time.Second)
			fmt.Println("aaa")
		}

	}
}

