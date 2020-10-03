package main

func main() {

}


func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices);i++  {
		dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
		dp[i][1] = max(dp[i - 1][1], - prices[i])
	}
	return dp[len(prices) - 1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}



func maxProfit(prices []int) int {
	maxProfit := 0
	if len(prices) < 2 {
		return maxProfit
	}
	minPrice := prices[0]
	for i := 1; i < len(prices); i ++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			tmp := prices[i] - minPrice
			if tmp > maxProfit {
				maxProfit = tmp
			}
		}
	}
	return maxProfit
}
