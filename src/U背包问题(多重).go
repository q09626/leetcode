package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n, &c)

	var v []int = make([]int, n+1)
	var w []int = make([]int, n+1)
	var m []int = make([]int, n+1)
	v[0] = 0
	w[0] = 0
	m[0] = 0

	for i := 1; i <= n; i++ {
		var vi, wi, mi int
		fmt.Scan(&vi, &wi, &mi)
		v[i] = vi
		w[i] = wi
		m[i] = mi
	}

	dp := make([][]int, n+1) // 长度加1
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, c+1) // 长度加1
	}

	// 填表
	for i := 1; i < n+1; i++ {
		for j := 1; j < c+1; j++ {
			// k=0已经包括了不放的情况
			for k := 0; k <= m[i] && k*w[i] <= j; k++ {
				// 每次循环和自己比较
				// k=0时，下面的比较包含了默认不放情况的值
				if dp[i][j] < dp[i-1][j-k*w[i]]+k*v[i] {
					dp[i][j] = dp[i-1][j-k*w[i]] + k*v[i]
				}
			}
		}
	}

	fmt.Println(dp[n][c])
}
