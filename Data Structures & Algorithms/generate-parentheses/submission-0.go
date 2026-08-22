func generateParenthesis(n int) []string {
	var res []string
	var path string
	var dfs func(int, int)

	dfs = func(open, closed int) {
		if open == n && closed == n {
			res = append(res, path)
			return
		}

		if open < n {
			path = path + "("
			dfs(open+1, closed)
			path = path[:len(path)-1]
		}
		
		if closed < open {
			path = path + ")"
			dfs(open, closed+1)
			path = path[:len(path)-1]
		}
	}

	dfs(0,0)

	return res
}
