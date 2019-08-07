/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	isCycle := false
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			isCycle = true
			break
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	if !isCycle {
		return nil
	}

	slow = head
	fast = fast.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}