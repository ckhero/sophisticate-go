package reverse_list

/**
https://leetcode-cn.com/problems/reverse-linked-list/
翻转链表
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var res *ListNode = nil
	for head != nil {
		curr := &ListNode{Val: head.Val}
		head = head.Next
		curr.Next = res
		res = curr
	}
	return res
}

func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next ==nil {
		return head
	}
	pre := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return pre
}