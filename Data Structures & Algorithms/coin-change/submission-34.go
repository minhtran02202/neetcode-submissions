// at ith coin:
// - take coin
// 	- remainder = amount - ith coint
// 	- coin count + 1
// - skip coin
// 	- remainder unchange
// 	- coint count unchance

// return min combinations of take and skip

// return min(1 + dfs(amount - coins[i], i), dfs(amount, i+1))



func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := 1; i < len(dp); i++{ dp[i] = amount + 1 }

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a-c >= 0 { dp[a] = min(dp[a], 1 + dp[a-c]) }
		}
	}

	if dp[amount] > amount { return -1 } else { return dp[amount] }
}
