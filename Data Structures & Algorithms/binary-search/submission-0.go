func search(nums []int, target int) int {
	n := len(nums)

	l, r := 0, n - 1

	for l <= r {
		m := (l + r) / 2

		if nums[m] == target { return m }

		if nums[m] > target {
			r = m - 1
		} else{
			l = m + 1
		}
	}

	return -1
}
