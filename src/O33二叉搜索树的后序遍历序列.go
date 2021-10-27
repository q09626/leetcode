package main

func verifyPostorder(postorder []int) bool {
	var post func(i, j int) bool
	post = func(i, j int) bool {
		if i >= j {
			return true
		}
		var a int
		for a = i; a < j; a++ {
			if postorder[a] >= postorder[j] {
				break
			}
		}
		var b int
		for b = a; b < j; b++ {
			if postorder[b] < postorder[j] {
				break
			}
		}
		return b == j && post(i, a-1) && post(a, j-1)
	}
	return post(0, len(postorder)-1)
}
