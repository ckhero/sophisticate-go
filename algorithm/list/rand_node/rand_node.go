package rand_node

import (
	"math/rand"
	"time"
)

/**
https://leetcode-cn.com/problems/linked-list-random-node/
382. 链表随机节点
蓄水池抽样
 */

type ListNode struct {
	Val int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	r *rand.Rand
}


/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	res := -1
	c := 1
	l := this.head
	for l != nil {
		r := this.r.Intn(c)
		if r == c - 1 {
			res = l.Val
		}
		c++
		l = l.Next
	}
	return res
}