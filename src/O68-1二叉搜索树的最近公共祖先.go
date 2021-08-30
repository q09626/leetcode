package main

func lowestCommonAncestor1(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor1(root.Right, p, q)
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor1(root.Left, p, q)
	}
	return root
}