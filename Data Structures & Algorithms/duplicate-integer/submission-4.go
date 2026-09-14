func hasDuplicate(nums []int) bool {
    h := make(map[int]bool, len(nums))

	for _, val := range nums {
		if h[val] { return true }
		h[val] = true
	}

	return false
}
