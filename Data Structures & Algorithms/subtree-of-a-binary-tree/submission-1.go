/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root==nil {return false}
	if subRoot==nil {return true}

	if track(root, subRoot) {return true}
	
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot) 
}

func track(r1, r2 *TreeNode) bool {
	if r1==nil || r2==nil {
		return r1==nil && r2==nil
	}

	if r1.Val != r2.Val {
		return false
	}

	return track(r1.Left, r2.Left) && track(r1.Right, r2.Right)
}