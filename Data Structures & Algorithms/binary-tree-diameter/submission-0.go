/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0 

	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {return 0}

		lheight := dfs(root.Left)
		rheight := dfs(root.Right)
		
		diameter = max(lheight+rheight, diameter)

		return max(lheight, rheight) + 1
	}

	dfs(root)
	return diameter
}
