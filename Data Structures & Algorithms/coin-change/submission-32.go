func coinChange(coins []int, amount int) int {
	minCoins := make([]int, amount+1)
	for i:=1; i<=amount; i++ {
		minCoins[i] = amount+1
	}

	for _, coin := range coins {
		for i:=1; i<=amount; i++ {
			if i - coin >= 0 {
				minCoins[i] = min(minCoins[i], 1+minCoins[i - coin])
			}
		}
	}

	if minCoins[amount] == amount+1 {
		return -1
	} else {
		return minCoins[amount]
	}
}
