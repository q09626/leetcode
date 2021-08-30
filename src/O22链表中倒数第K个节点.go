package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}