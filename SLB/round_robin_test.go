package SLB

import (
	"fmt"
	"testing"
)

func TestRoundRobin(t *testing.T) {
	nodes := make(map[int]string)
	nodes[0] = "0"
	nodes[1] = "1"
	nodes[2] = "2"
	count := map[string]int{}
	next := RoundRobin(nodes)
	for i := 0; i < 1000; i++ {
		count[next()]++
	}
	fmt.Println(count)
}