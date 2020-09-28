package rand_node

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	node1 := &ListNode{
		Val:  0,
	}
	node2 := &ListNode{
		Val:  1,
	}
	node3 := &ListNode{
		Val:  2,
	}
	node4 := &ListNode{
		Val:  3,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	obj := Constructor(node1)
	res := make([]int, 4)
	fmt.Println(res)

	for i := 0; i < 10000; i++ {
		res[obj.GetRandom()]++
	}
	fmt.Println(res)
}
