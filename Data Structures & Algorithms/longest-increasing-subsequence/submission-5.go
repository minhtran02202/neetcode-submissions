func lengthOfLIS(nums []int) int {
	finalRes := 0
	memo := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		res := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i]{
				res = max(res, 1 + memo[j])
			}
		}
		memo[i] = res
		finalRes = max(finalRes, res)
	}

	return finalRes
}
