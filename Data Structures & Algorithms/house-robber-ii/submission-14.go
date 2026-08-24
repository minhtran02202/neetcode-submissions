func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	return max(handle(nums[:n-1]), handle(nums[1:]))
}

func handle(nums []int) int {
	house1, house2 := 0, 0
	
	for _, val := range nums {
		house1, house2 = house2, max(house2, house1 + val)
	}

	return house2 
}