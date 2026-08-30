/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil { return false }

    var travel func(s, f *ListNode) bool
	travel = func(s, f *ListNode) bool {
		if s == nil || f == nil || f.Next == nil { return false }

		if s == f { return true }

		return travel(s.Next, f.Next.Next)
	}

	return travel(head, head.Next)
}
