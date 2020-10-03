package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/solution/
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

动态规划
f(i) 代表以i结尾的最大值
f(i)= max(i, f(i - 1) + i)
*/

func main()  {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	maxV := nums[0]
	for i := 1; i < len(nums); i ++ {
		nums[i] = max(nums[i], nums[i - 1] + nums[i])
		maxV = max(nums[i], maxV)
	}
	return maxV
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}