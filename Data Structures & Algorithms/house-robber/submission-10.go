

func rob(nums []int) int {
	res := 0
	robbed, skipped := 0, 0

	for _, val := range nums {
		res = max(val + robbed, skipped)
		robbed = skipped
		skipped = res
	}

	return res
}
