package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	start := dummy
	p := start.Next
	r := start.Next
	for {
		r = p // 新链表的开头等于合成后的尾
		for i := 1; i < k; i++ {
			if p == nil {
				break
			}
			p = p.Next
		}
		if p == nil {
			start.Next = r
			break
		}
		q := p.Next
		p.Next = nil
		start.Next = reverseList(r)
		start = r
		p = q
	}
	return dummy.Next
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	showList(reverseKGroup(head, 2))
}
