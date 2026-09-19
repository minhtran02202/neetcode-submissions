

import "slices"

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 { return 0 }

	slices.Sort(nums)
	local, res := 1, 1

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] { continue }
		if nums[i] - nums[i - 1] > 1 { local = 1; continue }
		local++
		res = max(local, res)
	}

	return res
}
