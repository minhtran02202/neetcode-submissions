func combine(n int, k int) [][]int {
	var result [][]int

	var dfs func(int, []int)
	dfs = func(i int, path []int) {
		if len(path) == k {
			t := append([]int{}, path...)
			result = append(result, t)
			return
		}

		for j := i; j <= n; j++ {
			dfs(j + 1, append(path, j))
		}
	}

	dfs(1, []int{})

	return result
}
