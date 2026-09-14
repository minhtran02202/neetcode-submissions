func productExceptSelf(nums []int) []int {
	prod := 1
	zcount := 0

	for _, num := range nums {
		if num != 0 { prod *= num } else { zcount++ }
	}

	res := make([]int, len(nums))

	if zcount > 1 { return res }

	for i, num := range nums {
		if zcount > 0 {
			if num == 0 {
				res[i] = prod
			} else {
				res[i] = 0
			}
		} else {
			res[i] = prod / num
		}
	}

	return res
}
