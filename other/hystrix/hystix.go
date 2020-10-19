package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"sync"
)

func main()  {
	hystrix.ConfigureCommand("test", hystrix.CommandConfig{
		Timeout:                5,
		MaxConcurrentRequests:  10,
		RequestVolumeThreshold: 1000,
		SleepWindow:            10,
		ErrorPercentThreshold:  50,
	})
	var wg sync.WaitGroup
	for i :=0; i < 100; i++ {
		wg.Add(1)
		defer wg.Done()

		go func() {
			hystrix.DoC(context.Background(), "test", func(ctx context.Context) error {
				// ...
				fmt.Println("succ")
				return nil
			}, func(i context.Context, e error) error {
				// ...
				fmt.Println("fail")
				return e
			})
		}()
	}
	wg.Wait()
}

