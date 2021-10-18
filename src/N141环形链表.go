package main

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
