func coinChange(coins []int, amount int) int {
	minCoins := make([]int, amount+1)

	for i:=1; i<=amount; i++ {
		minCoinsForThisAmount := amount+1

		for _, coin := range coins {
			if i - coin >= 0 && minCoins[i - coin] != -1 {
				minCoinsForThisAmount = min(minCoinsForThisAmount, 1 + minCoins[i - coin])
			}
		}

		if minCoinsForThisAmount == amount+1 {
			minCoins[i] = -1
		} else {
			minCoins[i] = minCoinsForThisAmount
		}
	}

	return minCoins[amount]
}
