/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil { return []int{} }
    var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		res = append(res, queue[len(queue)-1].Val)
		
		n := len(queue)
		for i:=0; i<n; i++ {
			pop := queue[0]
			queue = queue[1:]
			
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
	}

	return res
}
