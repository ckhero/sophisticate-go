package SLB

import (
	"math/rand"
	"time"
)

func Random(nodes map[int]string) string {
	rand.Int()
	r := rand.New(rand.NewSource(getRandomSeek()))
	index := r.Int() % len(nodes)
	return nodes[index]
}

func getRandomSeek() int64 {
	return time.Now().UnixNano()
}