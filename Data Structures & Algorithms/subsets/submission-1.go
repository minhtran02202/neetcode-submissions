func subsets(nums []int) [][]int {
	var res [][]int

	var backtrack func(i int, path []int)
	backtrack = func(i int, path []int) {
		if i == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		// choices: add or not
		
		// add
		path = append(path, nums[i])
		backtrack(i+1, path)

		path = path[:len(path)-1]

		// not
		backtrack(i+1, path)

	}

	backtrack(0, []int{})

	return res
}
