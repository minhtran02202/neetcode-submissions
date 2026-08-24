func coinChange(coins []int, amount int) int {
	cache := make(map[int]int)
	cache[0] = 0


	var minCoinsAt_toMake func(int) int
	minCoinsAt_toMake = func(amount int) int {
		if amount == 0 {
			return 0
		}

		if val, ok := cache[amount]; ok {
			return val
		}

		minCoinsForThisAmount := math.MaxInt

		for _, coin := range coins {
			diff := amount - coin

			if diff < 0 {
				continue
			}

			res := minCoinsAt_toMake(diff)
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

	return minCoinsAt_toMake(amount)
}
