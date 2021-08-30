package main

func deleteNode(head *ListNode, val int) *ListNode {
	var p, q *ListNode
	p, q = head, nil
	for p != nil && p.Val != val {
		q = p
		p = p.Next
	}
	if p != nil {
		if q == nil {
			return p.Next
		} else {
			q.Next = p.Next
			return head
		}
	}
	return head
}

func main() {
	head := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	showList(deleteNode(head, 0))
}