func climbStairs(n int) int {
	memo := make(map[int]int)
	
	var climb func(int) int

	climb = func(n int) int {
		if n == 1 || n == 2 {
			return n
		}
		if val, ok := memo[n]; ok {
			return val
		}

		memo[n] = climb(n-1) + climb(n-2)
		return memo[n]
	}

	return climb(n)
}
