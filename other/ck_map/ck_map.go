package ck_map

import (
	"fmt"
	"sort"
)

var keys []string

func SortRange(data map[string]interface{})  {

	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v)
	}
}