package SLB

import (
	"math/rand"
	"sync"
)

type Next func() string

func RoundRobin(nodes map[int]string) Next {
	var i = rand.Int()
	var mtx sync.Mutex
	return func() string {
		mtx.Lock()
		node := nodes[i % len(nodes)]
		i++
		mtx.Unlock()
		return node
	}
}
