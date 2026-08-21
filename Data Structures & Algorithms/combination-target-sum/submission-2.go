func combinationSum(nums []int, target int) [][]int {
    var res [][]int

	var dfs func([]int, []int, int)

	dfs = func(candidates, path []int, target int) {
		if target == 0 {
			t := append([]int{}, path...)
			res = append(res, t)
		}

		for i, val := range candidates {
			if target - val < 0 {
				continue
			}
			path = append(path, val)
			dfs(candidates[i:], path, target-val)
			path = path[:len(path)-1]
		}
	}

	dfs(nums, []int{}, target)

	return res
}
