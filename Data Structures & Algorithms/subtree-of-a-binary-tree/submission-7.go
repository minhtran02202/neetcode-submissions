/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return root != nil && (search(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot))
}

func search(r, sr *TreeNode) bool {
	if r == nil || sr == nil { return r == nil && sr == nil }
	return r.Val == sr.Val && search(r.Left, sr.Left) && search(r.Right, sr.Right)
}
