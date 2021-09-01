package main

import (
	"container/list"
	"fmt"
)

func showNums(nums []int) {
	for _, v := range nums {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func showList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func showTree(root *TreeNode) {
	var ans []int
	p := []*TreeNode{root}
	for len(p) > 0 {
		var q []*TreeNode
		for _, v := range p {
			ans = append(ans, v.Val)
			if v.Val == 0 {
				continue
			}
			if v.Left != nil {
				q = append(q, v.Left)
			} else {
				q = append(q, &TreeNode{0, nil, nil})
			}
			if v.Right != nil {
				q = append(q, v.Right)
			} else {
				q = append(q, &TreeNode{0, nil, nil})
			}
		}
		p = q
	}

	i := len(ans) - 1
	for i >= 0 {
		if ans[i] != 0 {
			break
		}
		i--
	}
	ans = ans[:i+1]
	showNums(ans)
}

type node struct {
	val int
	id  int
}

type nodes []node

func showListStruct(l list.List) {
	fmt.Println("val  id")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(node).val, e.Value.(node).id)
	}

	//for e := l.Back(); e != nil; e = e.Prev() {
	//	fmt.Println(e.Value.(node).val, e.Value.(node).id)
	//}
}
