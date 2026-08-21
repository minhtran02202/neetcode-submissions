/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root==nil{return true}

	diff := abs(height(root.Left) - height(root.Right))

	if diff > 1{
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(r *TreeNode) int {
	if r==nil{return 0}

	return max(height(r.Left), height(r.Right)) + 1 
}

func abs(x int) int{
	if x<0{return -x}
	return x
}
