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
	r := NewRoundRobin()
	for i := 0; i < 10; i++ {
		node := r.RoundRobin(nodes)

		fmt.Println(node())
	}
}