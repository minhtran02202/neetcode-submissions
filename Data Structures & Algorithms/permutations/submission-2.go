// each res item is len of nums
// diff order is a valid res item
// can't reuse nums[i] if nums[i] already considered for result

func permute(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	visited := make([]bool, n)
	var dfs func([]int)

	dfs = func(path []int) {
		if len(path) == n {
			res = append(res, append([]int{}, path...))
			return
		}

		for j := 0; j < n; j++ {
			if visited[j] { continue }
			visited[j] = true
			dfs(append(path, nums[j]))
			visited[j] = false
		}
	}

	dfs([]int{})

	return res
}
