package main

import "container/list"

func levelOrder3(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	s := list.New()
	s.PushBack(root)
	var p *TreeNode
	flag := true // 前面取，后面放，先放左，后方右
	for s.Len() > 0 {
		var tmp []int
		l := s.Len()
		for i := 0; i < l; i++ {
			if flag {
				p = s.Front().Value.(*TreeNode)
				s.Remove(s.Front())
				if p.Left != nil {
					s.PushBack(p.Left)
				}
				if p.Right != nil {
					s.PushBack(p.Right)
				}
			} else {
				p = s.Back().Value.(*TreeNode)
				s.Remove(s.Back())
				if p.Right != nil {
					s.PushFront(p.Right)
				}
				if p.Left != nil {
					s.PushFront(p.Left)
				}
			}
			tmp = append(tmp, p.Val)
		}
		flag = !flag
		ret = append(ret, tmp)
	}
	return ret
}
