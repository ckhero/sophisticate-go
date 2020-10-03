package main

/**
https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
摩尔投票法
 */
func majorityElement(nums []int) int {
	ans := 0
	count := 0
	for i  :=0; i < len(nums); i ++ {
		if count == 0 {
			ans = nums[i]
		}
		if ans == nums[i] {
			count ++
		} else {
			count --
		}
	}
	return ans
}
