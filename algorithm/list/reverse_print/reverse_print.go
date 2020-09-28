package reverse_print

/**
https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func ReversePrint(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	for i, j := 0, len(res) - 1; i < j; {
		res[i], res[j] = res[j], res[i]
		j--
		i++
	}
	return res
}