package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
)

func TestSortMap(t *testing.T) {
	m := map[string]string{
		"1":"222",
		"2":"2212",
		"3":"2232",
	}
	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(m[v])
	}
}

func TestSortedMap(t *testing.T) {
	m := NewMapList()
	var a,b,c Keyer
	a = &Element{
		key:   "aaa",
		value: "aavvv",
	}

	b = &Element{
		key:   "bbbb",
		value: "bbvvv",
	}

	c = &Element{
		key:   "ccc",
		value: "cccvvv",
	}
	m.Push(a)
	m.Push(b)
	m.Push(c)

	cb := func(data Keyer) {
		fmt.Print(data.GetVal())
		fmt.Println(data.GetKey())
	}
	m.Walk(cb)
}


func TestKeyType(t *testing.T) {
	m := make(map[KeyStruct]int)
	key1 := KeyStruct{a:1}
	m[key1] = 1
	fmt.Println(m)
}

func TestCopy(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 111
	data,_ := json.Marshal(m)
	m2 := make(map[string]int)
	_ = json.Unmarshal(data, &m2)
	fmt.Println(m2)
	m2["a"] = 2222
	fmt.Println(m)
	fmt.Println(m2)
}