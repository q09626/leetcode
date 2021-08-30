package main

func mirrorTree(root *TreeNode) *TreeNode {
	var mirrow func(root *TreeNode)
	mirrow = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		mirrow(root.Left)
		mirrow(root.Right)
	}

	mirrow(root)
	return root
}