

func combinationSum(nums []int, target int) [][]int {	
	var res [][]int
	var dfs func(int, int, []int)

	dfs = func(i, diff int, path []int) {
		if diff == 0 {
			res = append(res, append([]int{}, path...))
			return
		} 
		
		if diff < 0 || i == len(nums) { return }

		for j := i; j < len(nums); j++ {
			dfs(j, diff - nums[j], append(path, nums[j]))
		}
	}

	dfs(0, target, []int{})

	return res
}
