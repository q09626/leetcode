package main

func maxDepth(root *TreeNode) int {
	var max int = 0
	var core func(root *TreeNode, dep int)
	core = func(root *TreeNode, dep int){
		if root.Left == nil && root.Right == nil {
			dep++
			if dep > max {
				max = dep
			}
			return
		}
		if root.Left != nil {
			core(root.Left, dep+1)
		}
		if root.Right != nil {
			core(root.Right, dep+1)
		}
	}

	if root == nil {
		return 0
	}
	core(root, 0)
	return max
}