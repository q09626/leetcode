package main

func isSymmetric(root *TreeNode) bool {
	var leftAndRight func(left *TreeNode, right *TreeNode) bool
	leftAndRight = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return leftAndRight(left.Left, right.Right) && leftAndRight(left.Right, right.Left)
	}

	return leftAndRight(root, root)
}