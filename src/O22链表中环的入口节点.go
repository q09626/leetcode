package main

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next

		if fast == slow {
			p := head
			for fast != p {
				p = p.Next
				fast = fast.Next
			}
			return p
		}
	}
	return nil
}
