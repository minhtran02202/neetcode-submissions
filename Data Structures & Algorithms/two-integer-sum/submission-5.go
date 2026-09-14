func twoSum(nums []int, target int) []int {
    h := make(map[int]int, len(nums))

	for i, val := range nums {
		if ans, ok := h[target - val]; ok { return []int{ans, i} }
		
		h[val] = i
	}
	
	return []int{}
}
