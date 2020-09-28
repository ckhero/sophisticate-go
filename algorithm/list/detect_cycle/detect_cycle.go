package detect_cycle

/**
https://leetcode-cn.com/problems/linked-list-cycle-ii/s
 */
type ListNode struct {
	Val int
	Next *ListNode
}


func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	intersectionNode := getIntersectionNode(head)
	if intersectionNode == nil {
		return nil
	}
	for intersectionNode != head {
		head = head.Next
		intersectionNode = intersectionNode.Next
}
	return head
}

func getIntersectionNode(head *ListNode) *ListNode {
	l1, l2 := head, head
	for {
		if l1 ==  nil || l2 == nil || l2.Next == nil {
			return nil
		}
		l1 = l1.Next
		l2 = l2.Next.Next
		if l1 == l2 {
			break
		}
	}
	return l1
}