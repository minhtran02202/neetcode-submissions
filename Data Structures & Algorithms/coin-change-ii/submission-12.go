func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

    cache := make([][]int, len(coins))

	for i := range cache{
		cache[i] = make([]int, amount+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}	

	var combinationsOfCoinAt_toMake func (int, int) int
	combinationsOfCoinAt_toMake = func(at, target int) int {
		if target == 0 { return 1 }
		if at >= len(coins) || target < 0 { return 0 }

		if cache[at][target] != -1 { return cache[at][target] }

		cache[at][target] = combinationsOfCoinAt_toMake(at+1, target) + combinationsOfCoinAt_toMake(at, target-coins[at])

		return cache[at][target]
	}

	return combinationsOfCoinAt_toMake(0, amount)
}
