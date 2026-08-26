func cloneGraph(node *Node) *Node {
    visited := map[int]*Node{}

	var copy func(*Node) *Node
	copy = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		if c, ok := visited[n.Val]; ok {
			return c
		}

		c := &Node{
			Val: n.Val,
			Neighbors: make([]*Node, 0, len(n.Neighbors)),
		}

		visited[n.Val] = c

		for _, nei := range n.Neighbors {
			c.Neighbors = append(c.Neighbors, copy(nei))
		}

		return c
	}

	return copy(node)
}
