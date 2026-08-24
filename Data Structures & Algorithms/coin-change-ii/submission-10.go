func change(amount int, coins []int) int {
	dp := append([]int{1}, make([]int, amount)...)

	for _, c := range coins{
		for a := c; a<amount+1; a++{
			dp[a] += dp[a-c]
		}
	} 
	return dp[amount]
}
