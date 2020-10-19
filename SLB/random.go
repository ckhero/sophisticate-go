package SLB

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("SLB")
}
func Random(nodes map[int]string) string {
	rand.Int()
	r := rand.New(rand.NewSource(getRandomSeek()))
	index := r.Int() % len(nodes)
	return nodes[index]
}

func getRandomSeek() int64 {
	return time.Now().UnixNano()
}

func RandomWeight(nodes map[int]map[string]interface{}) Next {
	tw := getTotalWeight(nodes)
	return func() string {
		randomPos := rand.Intn(tw) + 1
		for _, v := range nodes {
			weight, _ := v["weight"].(int)

			if randomPos -= weight; randomPos <= 0 {
				ip, _ := v["ip"].(string)
				return ip
			}
		}
		return ""
	}
}

func getTotalWeight(nodes map[int]map[string]interface{}) int {
	tw := 0
	for _, v := range nodes {
		weight, _ := v["weight"].(int)
		tw += weight
	}
	return tw
}