package stock_three

func maxProfit(prices []int) int {
	maxK := 2
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2][2]int, len(prices))
	for i := 0; i < len(prices);i++  {

		for k := 1; k <= maxK; k ++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i - 1][k][0], dp[i - 1][k][1] + prices[i])
			dp[i][k][1] = max(dp[i - 1][k][1], dp[i - 1][k-1][0] - prices[i])
		}
	}
	return dp[len(prices) - 1][maxK][0]
}


func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}