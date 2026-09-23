// distinct ints
// find all unique combs of nums == target
// same nums can be reused
// target is positive int

// backtracking
// need: sorted array

// base conditions
// target == 0; found 1 path (not taking any nums)
// target < 0; over calc -> cur path not valid -> return 0
// if i == len nums

// actions: take cur num or take next num 

// import "slices"

func combinationSum(nums []int, target int) [][]int {
	// slices.Sort(nums)
	
	n := len(nums)
	var res [][]int
	var dfs func(int, int, []int)

	dfs = func(i, diff int, path []int) {
		if diff == 0 {
			res = append(res, append([]int{}, path...))
			return
		} else if diff < 0 || i == n { return }

		for j := i; j < n; j++ {
			dfs(j, diff - nums[j], append(path, nums[j]))
		}
	}

	dfs(0, target, []int{})

	return res
}
