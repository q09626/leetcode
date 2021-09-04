package main

import "fmt"

func main() {
	//n, c := 5, 10
	//v := []int{0, 8, 10, 4, 5, 5} // 补0，下标i=第i个
	//w := []int{0, 6, 4, 2, 4, 3}

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

	dp := make([][]int, n+1) // 长度加1
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, c+1)
	}

	// 初始化
	for i := 0; i < n+1; i++ {
		dp[i][0] = 0
	}
	for j := 0; j < c+1; j++ {
		dp[0][j] = 0
	}

	// 填表
	for i := 1; i < n+1; i++ {
		for j := 1; j < c+1; j++ {
			// 默认是放0次的情况
			dp[i][j] = dp[i-1][j]
			// 放0次的情况
			if w[i] > j {
				continue
			}
			// 放1次的情况
			if dp[i-1][j] < dp[i-1][j-w[i]]+v[i] {
				dp[i][j] = dp[i-1][j-w[i]] + v[i]
			}
		}
	}

	fmt.Println(dp[n][c])
}
