/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var trav func(*ListNode) *ListNode
	trav = func(node *ListNode) *ListNode {
		if node == nil { return nil }

		node.Next = trav(node.Next)

		n--
		if n == 0 { return node.Next } else { return node } 
	}

	return trav(head)
}
