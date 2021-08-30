package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dum := &ListNode{-1, nil}
	var p1, p2, cur *ListNode
	p1, p2, cur = l1, l2, dum
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			cur.Next = p1
			p1 = p1.Next
			cur = cur.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
			cur = cur.Next
		}
	}
	if p1 != nil {
		cur.Next = p1
	} else {
		cur.Next = p2
	}
	return dum.Next
}