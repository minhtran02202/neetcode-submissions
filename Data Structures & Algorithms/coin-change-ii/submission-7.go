func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, c := range coins{
		for x := c; x<amount+1; x++{
			dp[x] += dp[x-c]
		}
	} 
	return dp[amount]
}
