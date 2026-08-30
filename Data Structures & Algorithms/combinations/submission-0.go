func combine(n int, k int) [][]int {
	var result [][]int
	path := []int{}

	var dfs func(int)
	dfs = func(i int) {
		if len(path) == k {
			t := append([]int{}, path...)
			result = append(result, t)
			return
		}

		for j := i; j <= n; j++ {
			path = append(path, j)
			dfs(j + 1)
			path = path[: len(path) - 1]
		}
	}

	dfs(1)

	return result
}
