package main

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

func main() {
	root := &TreeNode{0, &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, nil, nil}}
	showNums(inorderTraversal(root))
}