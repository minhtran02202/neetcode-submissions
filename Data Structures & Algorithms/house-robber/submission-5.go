func rob(nums []int) int {
	res := make([]int, len(nums))
	
	for i := len(nums) - 1; i >= 0; i-- {
		skip := 0
		rob := nums[i]
		if i+1 < len(res) {
			skip = res[i+1]
		}

		if i+2 < len(res) {
			rob += res[i+2]
		}

		res[i] = max(skip, rob)
	}

	return res[0]
}
