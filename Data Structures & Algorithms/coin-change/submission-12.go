func coinChange(coins []int, amount int) int {
	cache := make(map[int]int)
	cache[0] = 0
	minCoins := make([]int, amount+1)

	for i:=1; i<=amount; i++ {
		if val, ok := cache[i]; ok {
			minCoins[i] = val
			continue
		}

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
			cache[i] = -1
		} else {
			cache[i] = minCoinsForThisAmount
		}

		minCoins[i] = cache[i]
	}

	// 6. Call the recursive function to kick off the process
	return minCoins[amount]
}
