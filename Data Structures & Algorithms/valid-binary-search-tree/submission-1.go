/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(r *TreeNode) bool {
    if r == nil { return true }

	return validate(r, math.MinInt, math.MaxInt)
}

func validate(r *TreeNode, l, u int) bool {
	if r == nil { return true }

	if l < r.Val && r.Val < u {
		return validate(r.Left, l, r.Val) && validate(r.Right, r.Val, u)
	}

	return false
}