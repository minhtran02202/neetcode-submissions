/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	if search(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func search(r, sr *TreeNode) bool {
	if r == nil || sr == nil {
		return r == nil && sr == nil
	}

	if r.Val != sr.Val { return false }

	return search(r.Left, sr.Left) && search(r.Right, sr.Right)
}
