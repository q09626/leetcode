package main

import "container/list"

type MinStack struct {
	a *list.List
	b *list.List
}

func Constructor30() MinStack {
	return MinStack{&list.List{}, &list.List{}}
}

func (m *MinStack) Push(x int)  {
	m.a.PushFront(x)
	if m.b.Len() == 0 || x <= m.b.Front().Value.(int) { // 相等时也要放进去
		m.b.PushFront(x)
	}
}

func (m *MinStack) Pop()  {
	x := m.a.Remove(m.a.Front()).(int)
	if m.b.Front().Value.(int) == x {
		m.b.Remove(m.b.Front())
	}
}

func (m *MinStack) Top() int {
	return m.a.Front().Value.(int)
}

func (m *MinStack) Min() int {
	return m.b.Front().Value.(int)
}