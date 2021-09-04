package main

import "fmt"

func main() {
	//n, c := 4, 10
	//v := []int{0,2,3,2,3} // 补0，下标i=第i个
	//w := []int{0,3,4,2,5}
	//m := []int{0,2,2,1,4}

	n, c := 3, 15
	v := []int{0, 2, 3, 4} // 补0，下标i=第i个
	w := []int{0, 3, 4, 5}
	m := []int{0, 4, 3, 2}

	dp := make([][]int, n+1) // 长度加1
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, c+1)
	}

	// 初始化
	//for i := 0; i < n+1; i++ {
	//	dp[i][0] = 0
	//}
	//for j := 0; j < c+1; j++ {
	//	dp[0][j] = 0
	//}

	// 填表
	for i := 1; i < n+1; i++ {
		for j := 1; j < c+1; j++ {
			// k=0已经包括了不放的情况，所以表中每个数据都有值
			for k := 0; k <= m[i] && k*w[i] <= j; k++ {
				// 每次循环和自己比较
				if dp[i][j] < dp[i-1][j-k*w[i]]+k*v[i] {
					dp[i][j] = dp[i-1][j-k*w[i]] + k*v[i]
				}
			}
		}
	}

	fmt.Println(dp[n][c])
}
