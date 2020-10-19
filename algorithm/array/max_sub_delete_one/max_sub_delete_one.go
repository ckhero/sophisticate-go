package main

import "math"

/**
https://leetcode-cn.com/problems/maximum-subarray-sum-with-one-deletion/
1186. 删除一次得到子数组最大和
给你一个整数数组，返回它的某个 非空 子数组（连续元素）在执行一次可选的删除操作后，所能得到的最大元素总和。

换句话说，你可以从原数组中选出一个子数组，并可以决定要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，然后该子数组（剩下）的元素总和是所有子数组之中最大的。

注意，删除一个元素后，子数组 不能为空。

请看示例：

示例 1：

输入：arr = [1,-2,0,3]
输出：4
解释：我们可以选出 [1, -2, 0, 3]，然后删掉 -2，这样得到 [1, 0, 3]，和最大。
示例 2：

输入：arr = [1,-2,-2,3]
输出：3
解释：我们直接选出 [3]，这就是最大和。
转移方程
f(i) 代表以i结尾的不删除的最大的值
g(i)代表删除元素的最大值
f(i)  = max(f(i - 1), arr[i])
g(i)  = max(g(i - 1) + arr[i], f(i - 1))
res =  max(res, f(i), g(i))
 */


func maximumSum(arr []int) int {
	f := make([]int, len(arr))
	g := make([]int, len(arr))
	f[0] = arr[0]
	g[0] = -int(math.MaxInt32)
	res := f[0]
	for i := 1; i < len(arr); i ++ {
		if i == 1 {

		}
		f[i] = max(f[i - 1] + arr[i], arr[i])
		g[i] = max(g[i - 1] + arr[i], f[i - 1])
		res = multiMax(res, f[i], g[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func multiMax(x, y, z int) int {
	if x > y && x > z {
		return x
	}
	if y > x && y > z {
		return y
	}

	return z
}