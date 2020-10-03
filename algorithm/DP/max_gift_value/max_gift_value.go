package main

import "fmt"

/**
https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
剑指 Offer 47. 礼物的最大价值
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？



示例 1:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
 */
func main() {
	grid := [][]int{
		{1,2,5},
		{3,2,1},
	}
	fmt.Println(maxValue(grid))
}

func maxValue(grid [][]int) int {
	dp := initDp(grid)
	rows := len(grid)
	cols := len(grid[0])
	for row := 1; row < rows; row ++ {
		for col :=1; col < cols; col++ {
			dp[row][col]  = max(dp[row][col - 1], dp[row - 1][col]) + grid[row][col]
		}
	}
	return dp[rows - 1][cols - 1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func initDp(grid [][]int) [][]int {
	dp := make([][]int, len(grid))
	col := len(grid[0])
	for i := 0; i < len(grid); i ++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i ++ {
		dp[i][0]  = dp[i - 1][0] + grid[i][0]
	}

	for i := 1; i < col; i ++ {
		dp[0][i]  = dp[0][i - 1] + grid[0][i]
	}
	return dp
}