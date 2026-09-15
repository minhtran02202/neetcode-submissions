/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil { return false }
	f, s := head.Next, head
	for f != s {
		if f.Next == nil || f.Next.Next == nil { return false }
		f, s = f.Next.Next, s.Next
	}
	return true
}
