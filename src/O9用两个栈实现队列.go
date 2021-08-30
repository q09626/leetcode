package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{stack1: list.New(), stack2: list.New()}
}

func (c *CQueue) AppendTail(val int) {
	c.stack1.PushBack(val)
}

func (c *CQueue) DeleteHead() int {
	if c.stack2.Len() == 0 {
		for c.stack1.Len() > 0 {
			if c.stack1.Len() > 0 {
				c.stack2.PushBack(c.stack1.Remove(c.stack1.Back()))
			}
		}
	}
	if c.stack2.Len() != 0 {
		val := c.stack2.Back()
		c.stack2.Remove(val)
		return val.Value.(int)
	}
	return -1
}

func main() {
	obj := Constructor()
	obj.AppendTail(1)
	ret := obj.DeleteHead()
	fmt.Println(ret)
}