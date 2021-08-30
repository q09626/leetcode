package main

import (
	"container/list"
	"fmt"
)

func reversePrint(head *ListNode) []int {
	l := list.New()
	p := head
	for p != nil {
		l.PushFront(p.Val)
		p = p.Next
	}

	var ans []int
	for l.Len() != 0 {
		ans = append(ans, l.Front().Value.(int))
		l.Remove(l.Front())
	}
	return ans
}

func main() {
	head := &ListNode{1, &ListNode{2,&ListNode{3,&ListNode{4,&ListNode{5,nil}}}}}
	ans := reversePrint(head)
	for _, v := range ans {
		fmt.Printf("%d ", v)
	}
}