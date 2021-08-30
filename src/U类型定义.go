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
	Val int
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
	Val int
	Left *TreeNode
	Right *TreeNode
}

type node struct {
	val int
	id int
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