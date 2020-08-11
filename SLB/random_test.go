package SLB

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	nodes := make(map[int]string)
	nodes[0] = "0"
	nodes[1] = "1"
	nodes[2] = "2"
	res := make(map[string]int)
	for i :=0; i < 10; i++ {
		node := Random(nodes)
		if _, ok := res[node]; ok {
			res[node] += 1
		} else {
			res[node] = 1
		}
	}
	for k, v := range res {
		fmt.Printf("key = %s; volume =%d \n", k, v)
	}
}

func TestRandomWeight(t *testing.T) {
	nodes := make(map[int]map[string]interface{})
	nodes[0] = map[string]interface{}{
		"ip": "0",
		"weight": 5,
	}
	nodes[1] = map[string]interface{}{
		"ip": "1",
		"weight": 1,
	}
	nodes[2] = map[string]interface{}{
		"ip": "2",
		"weight": 4,
	}

	next := RandomWeight(nodes)
	count := map[string]int{}
	for i := 0; i < 10000; i++ {
		count[next()]++
	}
	fmt.Printf("%#v", count)
}