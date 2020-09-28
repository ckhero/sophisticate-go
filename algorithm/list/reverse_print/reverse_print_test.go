package reverse_print

import (
	"fmt"
	"testing"
)

func TestReversePrint(t *testing.T) {
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

	fmt.Println(ReversePrint(node1))
	fmt.Println(ReversePrint(node4))
}