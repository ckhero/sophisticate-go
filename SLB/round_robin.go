package SLB

import (
	"sync"
)
type RoundRobinBody struct {
	currIndex int
}

func NewRoundRobin() *RoundRobinBody {
	return &RoundRobinBody{
		currIndex: 1,
	}
}
func (r *RoundRobinBody) RoundRobin(nodes map[int]string) func() string {
	var mtx sync.Mutex
	return func() string {
		mtx.Lock()
		node := nodes[r.currIndex % len(nodes)]
		r.currIndex++
		mtx.Unlock()
		return node
	}
}
