func findTargetSumWays(nums []int, target int) int {
	cache := make([]map[int]int, len(nums) + 1)

	for i := 0; i <= len(nums); i++ {
		cache[i] = make(map[int]int)
	}

	cache[0][0] = 1

	for i := 0; i < len(nums); i++ {
		for total, count := range cache[i] {
			cache[i + 1][total + nums[i]] += count
			cache[i + 1][total - nums[i]] += count
		}
	}

	return cache[len(nums)][target]
}
