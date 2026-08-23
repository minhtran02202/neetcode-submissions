func coinChange(coins []int, amount int) int {
	cache := make(map[int]int)
	cache[0] = 0
	minCoins := make([]int, amount+1)

	var minCoinsAt_toMake func(int) int
	minCoinsAt_toMake = func(amount int) int {
		if val, ok := cache[amount]; ok {
			return val
		}

		minCoinsForThisAmount := math.MaxInt

		for _, coin := range coins {
			diff := amount - coin

			if diff < 0 {
				continue
			}

			res := minCoins[diff]
			if res != -1 {
				minCoinsForThisAmount = min(minCoinsForThisAmount, 1+res)
			}
		}

		if minCoinsForThisAmount == math.MaxInt {
			cache[amount] = -1
		} else {
			cache[amount] = minCoinsForThisAmount
		}

		return cache[amount]
	}

	for i:=1; i<=amount; i++ {
		minCoins[i] = minCoinsAt_toMake(i)
	}

	// 6. Call the recursive function to kick off the process
	return minCoins[amount]
}
