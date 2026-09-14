func twoSum(nums []int, target int) []int {
    h := make(map[int]int, len(nums))

	var res []int

	for i, val := range nums {
		if ans, ok := h[target - val]; ok {
			res = []int{ans, i}
			break
		}
		h[val] = i
	}
	
	return res
}
