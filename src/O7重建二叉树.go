package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	num := preorder[0]
	if len(preorder) == 1 {
		return &TreeNode{num, nil, nil}
	}
	i := 0
	for i < len(inorder) {
		if inorder[i] == num {
			break
		}
		i++
	}
	root := &TreeNode{num, buildTree(preorder[1:i+1], inorder[:i]), buildTree(preorder[i+1:], inorder[i+1:])}
	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	showTree(root)
}
