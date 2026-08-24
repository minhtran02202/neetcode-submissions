// given integer array coins and integer amount

// Return number of distinct combinations that total up to amount. If impossible, return 0.

// Assume unlimited number of each coin and that each value in coins is unique

// Input: amount = amount, coins = coins[]

// Output: 4

// with amount, try take one of each coins:
// 	amount -= coins[i]
// 	amount -= coins[i+1]
// 	amount -= coins[i+2]
// 	...

// observe: when take a coin, amount decrease by coins[i]
// meaning new amount is amount = amount - coins[i]
// meaning a way is found when amount == 0

// observe: distinct comb meaning two comb with same number of coins and number of each coins is not possible

// observe: aside from take, we can decide to skip that coin

// actions: take, skip

// take: comb = 1 + combination(amount - coins[i], i)

// skip: comb = combinations(amount, i+1)

// base case:
// 	amount 0: return 1 (1 way to achive 0 is skip action)
// 	if amount - coin[i] < 0 { return 0 } (invalid)



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

		if cache[at][target] != -1 {
			return cache[at][target]
		}

		cache[at][target] = combinationsOfCoinAt_toMake(at+1, target) + combinationsOfCoinAt_toMake(at, target-coins[at])

		return cache[at][target]
	}

	return combinationsOfCoinAt_toMake(0, amount)
}
