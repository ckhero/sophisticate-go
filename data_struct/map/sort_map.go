package main

import (
	"container/list"
	"context"
	"fmt"
)

type Keyer interface {
	GetKey() string
	GetVal() string
}

func main() {
	defer func() {
		if e := recover(); e !=nil {
			fmt.Println(e)
		}
	}()
	TestPanic1()
	//TestWRPanic()
}

func TestPanic1() {

	go TestPanic2()
	panic("22222")
}
func TestPanic2() {
	defer func() {
		if e := recover(); e !=nil {
			fmt.Println(2, e)
		}
	}()
	panic("22222")
}

func TestWRPanic()  {
	defer func() {
		if e := recover(); e !=nil {
			fmt.Println("catch it 2", e)
		}
	}()
	s := map[int]int{}
	for i :=0;  i < 20; i++ {
		go func(t int) {
			defer func() {
				if e := recover(); e !=nil {
					fmt.Println("catch it", e)
				}
			}()
			if t % 10 != 0 {
				s[0] = 2
			} else {
				_ = s[0]
			}
		}(i)
	}
}
type MapList struct {
	dataMap map[string]*list.Element
	dataList  *list.List
}

func NewMapList() *MapList {
	return &MapList{
		dataMap:  make(map[string]*list.Element),
		dataList: list.New(),
	}
}

func (m *MapList) Exists(data Keyer) bool {
	_, exists := m.dataMap[data.GetKey()]
	return exists
}

func (m *MapList) Push(data Keyer) bool {
	if m.Exists(data) {
		return false
	}
	ele := m.dataList.PushBack(data)
	m.dataMap[data.GetKey()] = ele
	return true
}

func (m *MapList) Remove(data Keyer) {
	if !m.Exists(data) {
		return
	}

	m.dataList.Remove(m.dataMap[data.GetKey()])
	delete(m.dataMap, data.GetKey())
}

func (m *MapList) Walk(cb func(data Keyer)) {
	for ele := m.dataList.Front(); ele != nil; ele = ele.Next() {
		cb(ele.Value.(Keyer))
	}
}

type Element struct {
	key string
	value string
}

func (e *Element) GetKey() string {
	return e.key
}

func (e *Element) GetVal() string {
return e.value
}

type KeyStruct struct {
	a int
}