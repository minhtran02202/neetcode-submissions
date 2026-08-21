/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    return search (root, root.Val)
}

func search(root *TreeNode, max int) int {
	if root == nil { return 0 }
	good := 0

	if root.Val >= max {
		good = 1
		max = root.Val
	}

	return good + search(root.Left, max) + search(root.Right, max)
}
