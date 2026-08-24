func coinChange(coins []int, amount int) int {
	minCoins := make([]int, amount+1)

	for i:=1; i<=amount; i++ {
		minCoinsForThisAmount := math.MaxInt

		for _, coin := range coins {
			diff := i - coin

			if diff < 0 {
				continue
			}

			res := minCoins[diff]
			if res != -1 {
				minCoinsForThisAmount = min(minCoinsForThisAmount, 1+res)
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
