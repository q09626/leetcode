package main

import "fmt"

func buildTree(preorder []int, inorder []int) *TreeNode {

	num := preorder[0]
	i := 0
	for i < len(inorder) {
		if inorder[i] == num {
			break
		} else {
			i++
		}
	}
	root := &TreeNode{num, buildTree(preorder[1:i+1], inorder[i+1:]), buildTree(preorder[:i], inorder[i+1:])}
	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root.Val)
}
