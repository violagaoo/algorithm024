package main

//322. 零钱兑换
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//初始值
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j >= coins[i-1] {
				//dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-coins[i-1]]+1))) //用时击败23.46%
				dp[j] = min(dp[j], dp[j-coins[i-1]]+1) //用时击败92.45%
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
