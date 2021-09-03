package main

// dfs
func movingCount1(m int, n int, k int) int {
	var compute func(i, j int) int
	compute = func(i, j int) int {
		ans := 0
		for i > 0 {
			ans += i % 10
			i /= 10
		}
		for j > 0 {
			ans += j % 10
			j /= 10
		}
		return ans
	}

	count := 0
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] == true || compute(i, j) > k {
			return
		}
		visited[i][j] = true
		count++
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	dfs(0, 0)
	return count
}

// bfs
func movingCount2(m int, n int, k int) int {
	var compute func(i, j int) int
	compute = func(i, j int) int {
		ans := 0
		for i > 0 {
			ans += i % 10
			i /= 10
		}
		for j > 0 {
			ans += j % 10
			j /= 10
		}
		return ans
	}

	count := 0
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// 向右和向下的方向数组
	dx := []int{0, 1}
	dy := []int{1, 0}
	// 保存某广度的对象
	q := make([][]int, 0)

	q = append(q, []int{0, 0})
	count++
	visited[0][0] = true
	for len(q) > 0 {
		p := make([][]int, 0)
		for i := 0; i < len(q); i++ {
			x, y := q[i][0], q[i][1]
			for j := 0; j < 2; j++ {
				tx := dx[j] + x
				ty := dy[j] + y
				if tx < 0 || tx >= m || ty < 0 || ty >= n || visited[tx][ty] == true || compute(tx, ty) > k {
					continue
				}
				p = append(p, []int{tx, ty})
				count++
				visited[tx][ty] = true
			}
		}
		q = p
	}
	return count
}
