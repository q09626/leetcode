package main

import "fmt"

func main() {
	n, c := 5, 10
	v := []int{0, 8, 10, 4, 5, 5} // 补0，下标i=第i个
	w := []int{0, 6, 4, 2, 4, 3}

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
			if w[i] > j {
				dp[i][j] = dp[i-1][j]
				continue
			}
			if dp[i-1][j] > dp[i-1][j-w[i]]+v[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-w[i]] + v[i]
			}
		}
	}

	fmt.Println(dp[n][c])
}
