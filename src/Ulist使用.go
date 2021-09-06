package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := list.List{} // 返回的是引用
	//stack := list.New() // 返回的是指针
	stack.PushBack(node{1, 1})
	stack.PushBack(node{4, 2})
	stack.PushBack(node{2, 3})
	showListStruct(stack)

	four := stack.PushFront(node{8, 4})
	showListStruct(stack)

	stack.InsertBefore(node{6, 5}, stack.Back())
	showListStruct(stack)

	stack.InsertAfter(node{9, 6}, stack.Front())
	showListStruct(stack)

	stack.MoveToBack(four)
	showListStruct(stack)

	stack.MoveBefore(stack.Front(), stack.Back())
	showListStruct(stack)

	for i := stack.Front(); i != nil; i = i.Next() {
		if i.Value.(node).id == 1 {
			stack.Remove(i)
			fmt.Println(i.Value.(node).id)
		}
	}
	//showListStruct(stack)
}
