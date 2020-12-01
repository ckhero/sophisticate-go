package main

import "fmt"

/**
https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/

 */
type myInt int
type myI = int
func main() {
	var a int = 2
	var b myI = a
	var c myInt = myInt(a)
	fmt.Println(a,b,c)
}
 type ListNode struct {
	 Val int
	 Next *ListNode
 }

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	l1 := headA
	_l2 := headB
	for l1 != _l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headA
		}

		if _l2 != nil {
			_l2 = _l2.Next
		} else {
			_l2 = headB
		}
	}
	return l1
}