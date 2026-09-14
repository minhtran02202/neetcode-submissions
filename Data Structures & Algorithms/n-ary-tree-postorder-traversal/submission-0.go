/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	res := []int{}

	var dfs func(*Node)
	dfs = func(n *Node) {
		if n == nil { return }

		for _, c := range n.Children {
			dfs(c)
		}

		res = append(res, n.Val)
	}

	dfs(root)

	return res
}
