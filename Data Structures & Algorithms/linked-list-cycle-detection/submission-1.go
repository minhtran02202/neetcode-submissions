/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil { return false }

	f, s := head.Next, head

	for f != nil {
		if f == s {
			return true
		}
		if f.Next == nil {
			return false
		}
		f = f.Next.Next
		s = s.Next
	}

	return false
}
