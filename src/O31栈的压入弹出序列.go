package main

import "container/list"

func validateStackSequences(pushed []int, popped []int) bool {
	s := list.New()
	i := 0
	for _, v := range pushed {
		s.PushBack(v)
		for s.Len() > 0 && s.Back().Value.(int) == popped[i] {
			i++
			s.Remove(s.Back())
		}
	}
	if s.Len() != 0 {
		return false
	}
	return true
}
