package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n, &c)

	var v []int = make([]int, n+1)
	var w []int = make([]int, n+1)
	v[0] = 0
	w[0] = 0

	for i := 1; i <= n; i++ {
		var vi, wi int
		fmt.Scan(&vi, &wi)
		v[i] = vi
		w[i] = wi
	}

	dp := make([]int, c+1) // 长度加1

	// 填表
	for i := 1; i < n+1; i++ {
		for j := c; j > 0; j-- {
			// 已经包含了默认不放值
			if w[i] > j {
				continue
			}
			if dp[j] < dp[j-w[i]]+v[i] {
				dp[j] = dp[j-w[i]] + v[i]
			}
		}
	}
	fmt.Println(dp[c])
}
