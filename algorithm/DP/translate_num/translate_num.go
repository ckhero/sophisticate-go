package main

import (
	"fmt"
	"strconv"
)

/**
https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
剑指 Offer 46. 把数字翻译成字符串
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。



示例 1:

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"


提示：

0 <= num < 231
通过次数49,373提交次数91,485
 */


func main() {
	//fmt.Println(translateNum(12258))
	fmt.Println(TranslateNum(1068385902))
	fmt.Println(TranslateNum2(1068385902))
}

func TranslateNum(num int) int {
	nums := strconv.Itoa(num)
	dp := make([]int, 1)
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		pre := nums[i - 1: i + 1]
		dp = append(dp, dp[ i - 1])
		if pre <= "25" && pre >= "10" {
			if i == 1 {
				dp[i] += 1
			} else {
				dp[i] += dp[ i - 2]
			}
		}
	}
	return dp[len(nums) - 1]
}

func TranslateNum2(num int) int {
	nums := strconv.Itoa(num)
	p1, p2, res :=0, 0, 1
	for i := 0; i < len(nums); i++ {
		res, p1, p2 = 0, res,  p1
		res = p1
		if i == 0 {
			continue
		}
		pre := nums[i - 1: i + 1]
		if pre <= "25" && pre >= "10" {
			res += p2
		}
	}
	return res
}

