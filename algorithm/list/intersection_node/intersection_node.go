package intersection_node

/**
https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/

 */


 type ListNode struct {
	 Val int
	 Next *ListNode
 }

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	l1 := headA
	l2 := headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headA
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headB
		}
	}
	return l1
}