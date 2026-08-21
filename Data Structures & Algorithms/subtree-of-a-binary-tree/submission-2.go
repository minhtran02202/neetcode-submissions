/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }

    if track(root, subRoot) {
        return true
    }

    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func track(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil || subRoot == nil {
        return root == nil && subRoot == nil
    }

    if root.Val != subRoot.Val {
        return false
    }

    return track(root.Left, subRoot.Left) && track(root.Right, subRoot.Right)
}