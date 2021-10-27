package main

func levelOrder2(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			tmp := q[j]
			ret[i] = append(ret[i], tmp.Val)
			if tmp.Left != nil {
				p = append(p, tmp.Left)
			}
			if tmp.Right != nil {
				p = append(p, tmp.Right)
			}
		}
		q = p
	}
	return ret
}
