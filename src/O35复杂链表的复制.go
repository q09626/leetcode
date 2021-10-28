package main

import "fmt"

func copyRandomList(head *RNode) *RNode {
	if head == nil {
		return nil
	}
	cur := head
	var tmp *RNode
	for cur != nil {
		tmp = &RNode{cur.Val, cur.Next, nil}
		cur.Next = tmp
		cur = tmp.Next
	}
	cur = head
	for cur != nil {
		tmp = cur.Next
		if cur.Random != nil {
			tmp.Random = cur.Random.Next
		} else {
			tmp.Random = nil
		}
		cur = tmp.Next
	}
	cur = head
	newHead := cur.Next
	for cur != nil {
		tmp = cur.Next
		cur.Next = tmp.Next
		if tmp.Next != nil {
			tmp.Next = tmp.Next.Next
		}
		cur = cur.Next
	}
	return newHead.Next
}

func PrintRNode(head *RNode) {
	cur := head
	for cur != nil {
		fmt.Print("(", cur.Val, ",")
		if cur.Random != nil {
			fmt.Print(cur.Random.Val, ") ")
		} else {
			fmt.Print(0, ") ")
		}
		cur = cur.Next
	}
	fmt.Println()
}

func main() {
	head := &RNode{7, &RNode{13, &RNode{11, &RNode{10, &RNode{1, nil, nil}, nil}, nil}, nil}, nil}
	head.Random = nil
	head.Next.Random = head
	head.Next.Next.Random = head.Next.Next.Next.Next
	head.Next.Next.Next.Random = head.Next.Next
	head.Next.Next.Next.Next.Random = head
	PrintRNode(head)
	PrintRNode(copyRandomList(head))
}
