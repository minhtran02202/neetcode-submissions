func twoSum(nums []int, target int) []int {
	p1, p2 := 0, len(nums)-1

	for p1 < p2 {
		if nums[p1] + nums[p2] == target {
			return []int{p1+1, p2+1}
		} else if nums[p1] + nums[p2] > target {
			p2--
		} else {
			p1++
		}
	}

	return []int{}
}
