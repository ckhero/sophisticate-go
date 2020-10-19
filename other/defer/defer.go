package main

import (
	"fmt"
	"os"
	"time"
)

func Test(i int)  {
	fmt.Println(i)
}
func main()  {
	Test4()
	fmt.Println(three())
	fmt.Println(two())
	i := 2
	defer Test(i)
	i = 3
	defer fmt.Println("defer main")
	var user = os.Getenv("USER_")

	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err: ", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}

			// 此处不会执行
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(100)
	fmt.Println("end of main function")
}


func one() {
	var whatever [3]struct{}

	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func two() ( int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Printf("two %d \n", t)
	}()
	return t
}

func three() (r int) {
	defer func(r int) {
		r = r + 5
		fmt.Printf("three%d \n", r)
	}(r)
	return 1
}

func Test4() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
}
func DeferFunc1(i int)(t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc3(i int)(t int) {
	defer func() {
		t += i
	}()
	return 2
}