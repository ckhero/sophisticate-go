	package has_cycle

/**
https://leetcode-cn.com/problems/linked-list-cycle/
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	l1, l2 := head, head
	for l2 != nil {
		l1 = l1.Next
		if l2.Next == nil {
			break
		}
		l2 = l2.Next.Next
		if l1 == l2 {
			return true
		}
	}
	return false
}