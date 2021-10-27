package main

import "fmt"

func pathSum(root *TreeNode, target int) [][]int {
	var ret [][]int
	var tmp []int // 有点bug，tmp加进ret后，之后的修改也会影响ret里面的值
	var path func(node *TreeNode, sum int)
	path = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		k := sum + node.Val
		if node.Left == nil && node.Right == nil {
			if k == target {
				tmp = append(tmp, node.Val)
				ret = append(ret, tmp)
				tmp = tmp[:len(tmp)-1]
			}
		} else {
			tmp = append(tmp, node.Val)
			path(node.Left, k)
			path(node.Right, k)
			tmp = tmp[:len(tmp)-1]
		}
	}
	path(root, 0)
	return ret
}

func main() {
	root := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, &TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}}}}
	ans := pathSum(root, 22)
	fmt.Println(len(ans))
}
