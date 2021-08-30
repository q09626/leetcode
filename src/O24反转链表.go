package main

func reverseList(head *ListNode) *ListNode {
	var p, q, r *ListNode
	p, q, r = head, nil, nil
	for p != nil {
		r = p.Next
		p.Next = q
		q = p
		p = r
	}
	return q
}