func rob(nums []int) int {
	n := len(nums)

	if n == 0 { return 0 }
	if n == 1 { return nums[0] }

    house1 := nums[0]
    house2 := max(nums[0], nums[1])
	
	for i := 2; i < n; i++ {
		t := house2
		house2 = max(house2, nums[i] + house1)
		house1 = t
	}

	return house2
}
