func validTree(n int, edges [][]int) bool {
    if len(edges) > n-1 {
		return false
	}

	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visit := make(map[int]bool)

	var dfs func(int, int) bool

	dfs = func(n, p int) bool {
		if visit[n] { return false }
		visit[n] = true

		for _, nei := range adj[n] {
			if nei == p { continue }
			
			if !dfs(nei, n) { return false }
		}

		return true
	}

	return dfs(0, -1) && len(visit) == n
}
