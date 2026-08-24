func rob(nums []int) int {
	memo := make([]int, len(nums))
	res := make([]int, len(nums))

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

		skip := 0
		rob := nums[i]
		if i+1 < len(res) {
			skip = res[i+1]
		}

		if i+2 < len(res) {
			rob += res[i+2]
		}

		memo[i] = max(skip, rob)

		return memo[i]
	}

	
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = robTotal_startAtHouse(i)
	}

	return res[0]
}
