func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(int, []int)

	dfs = func(i int, path []int) {
		if i == len(nums) { res = append(res, append([]int{}, path...)); return }

		dfs(i + 1, path); dfs(i + 1, append(path, nums[i]))
	}

	dfs(0, []int{})

	return res
}
