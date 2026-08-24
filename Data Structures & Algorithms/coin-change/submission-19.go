func coinChange(coins []int, amount int) int {
	minCoins := make([]int, amount+1)

	for i:=1; i<=amount; i++ {
		minCoinsForThisAmount := math.MaxInt

		for _, coin := range coins {
			if i - coin < 0 {
				continue
			}
			if minCoins[i - coin] != -1 {
				minCoinsForThisAmount = min(minCoinsForThisAmount, 1 + minCoins[i - coin])
			}
		}

		if minCoinsForThisAmount == math.MaxInt {
			minCoins[i] = -1
		} else {
			minCoins[i] = minCoinsForThisAmount
		}
	}

	// 6. Call the recursive function to kick off the process
	return minCoins[amount]
}
