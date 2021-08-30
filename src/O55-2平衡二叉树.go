package main

import "math"

func isBalanced(root *TreeNode) bool {
	var core func(root *TreeNode) (int, bool)
	core = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		l, isB1 := core(root.Left)
		if !isB1 {
			return -1, false
		}
		r, isB2 := core(root.Right)
		if !isB2 {
			return -1, false
		}
		if math.Abs(float64(l-r)) < 2 {
			return int(math.Max(float64(l), float64(r)))+1, true
		} else {
			return -1, false
		}
	}

	_, isB := core(root)
	if isB {
		return true
	} else {
		return false
	}
}