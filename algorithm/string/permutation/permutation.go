package main

import "fmt"

/**
https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
剑指 Offer 38. 字符串的排列
输入一个字符串，打印出该字符串中字符的所有排列。



你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。



示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]


限制：

1 <= s 的长度 <= 8

通过次数46,124提交次数84,784
 */

func main()  {
	fmt.Println(permutation("abc"))
}

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	resM := make(map[string]int)
	path := make([]byte, 0)
	helper([]byte(s), path, resM, 0)
	res := make([]string, 0)
	for k := range resM {
		if k != " " {
			res = append(res, k)
		}
	}
 	return res
}

func helper(data []byte, path []byte, resM map[string]int, count int)  {
	if count == len(data) {
		resM[string(path)] = 0
		return
	}
	for k, v := range data {
		if v == '/' {
			continue
		}
		data[k] = '/'
		path = append(path, v)
		helper(data, path, resM, count + 1)
		path = path[0:count]
		data[k] = v
	}
}