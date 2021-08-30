package main

func kthLargest(root *TreeNode, k int) int {
	count := k
	var ans int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Right)
		if count == 0 {
			return
		}
		count--
		if count == 0 {
			ans = root.Val
			return
		}
		inorder(root.Left)
	}

	inorder(root)
	return ans
}