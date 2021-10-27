package main

import "container/list"

func levelOrder1(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return []int{}
	}
	s := list.New()
	s.PushBack(root)
	for s.Len() > 0 {
		tmp := s.Front().Value.(*TreeNode)
		s.Remove(s.Front())
		ret = append(ret, tmp.Val)
		if tmp.Left != nil {
			s.PushBack(tmp.Left)
		}
		if tmp.Right != nil {
			s.PushBack(tmp.Right)
		}
	}
	return ret
}
