/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  go to end
//  from end, n-- until n==0
//  if n==0, remove this node
//  remove by return next node instead of this node back to prev.Next

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var trav func(*ListNode) *ListNode

	trav = func(node *ListNode) *ListNode {
		if node == nil { return nil }

		node.Next = trav(node.Next)

		n--
		if n == 0 {
			return node.Next
		}
		return node
	}

	return trav(head)
}
