package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := nodes{node{1,1}, node{4,2}, node{2,3}}
	nums.Show()
	sort.Sort(nums)
	nums.Show()
}

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Less(i, j int) bool {
	return n[i].val > n[j].val
}

func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n nodes) Show() {
	fmt.Println("val id")
	for _, v := range n {
		fmt.Println(v.val, v.id)
	}
}