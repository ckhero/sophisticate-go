package main

import "fmt"

func main() {
	s := make(map[string]string)
	ss := new(map[string]string)
	s["aa"] = "a"
	fmt.Println(*ss)
}
