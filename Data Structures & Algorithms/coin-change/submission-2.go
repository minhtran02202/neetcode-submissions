func coinChange(coins []int, amount int) int {
	cache := make(map[int]int)
	cache[0] = 0

	// 1. Initialize the rest of the cache with -1 to mean "unvisited"
	// for i := 1; i < len(cache); i++ {
	// 	cache[i] = -1 
	// }

	var minCoinsAt_toMake func(int) int
	minCoinsAt_toMake = func(amount int) int {
		// 2. Base Case: If amount is 0, it takes 0 coins
		if amount == 0 {
			return 0
		}

		// 3. Memo step: Trigger if it's NOT -1
		if val, ok := cache[amount]; ok {
			return val
		}

		// 4. Temporary variable to find the minimum for the current amount
		minCoinsForThisAmount := math.MaxInt

		// choices: try each coin
		for _, coin := range coins {
			diff := amount - coin

			if diff < 0 {
				continue
			}

			res := minCoinsAt_toMake(diff)
			// Only update if the sub-problem actually found a valid combination
			if res != -1 {
				minCoinsForThisAmount = min(minCoinsForThisAmount, 1+res)
			}
		}

		// 5. If it stays MaxInt, it means this amount is impossible
		if minCoinsForThisAmount == math.MaxInt {
			cache[amount] = -1
		} else {
			cache[amount] = minCoinsForThisAmount
		}

		return cache[amount]
	}

	// 6. Call the recursive function to kick off the process
	return minCoinsAt_toMake(amount)
}
