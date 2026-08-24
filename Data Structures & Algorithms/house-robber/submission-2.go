func rob(nums []int) int {
	memo := make([]int, len(nums))

	for i := 0; i < len(memo); i++{
		memo[i] = -1
	}

    var robTotal_startAtHouse func (int) int
	robTotal_startAtHouse = func(i int) int {
		if i >= len(nums) {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		memo[i] = max(robTotal_startAtHouse(i+1), nums[i] + robTotal_startAtHouse(i+2))

		return memo[i]
	}

	return robTotal_startAtHouse(0)
}
