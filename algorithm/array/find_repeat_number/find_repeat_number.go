package find_repeat_number


/**
https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
 */

/**
正常排序下，如果没重复的，i == nums[i],如果i！=nums[i]的话，但是 nums[nums[i]] == nums[i]的话说明是重复的
 */
func findRepeatNumberOne(nums []int) int {
	for i :=0; i < len(nums); i ++ {
		if i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			}
		}
	}
	return -1
}