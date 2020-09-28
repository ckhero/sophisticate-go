package _map

import (
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