package main

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	ans := ""
	for l := 1; l <= len(s); l++ {
		for i := 0; i+l-1 < len(s); i++ {
			j := i + l - 1
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				if l > len(ans) {
					ans = s[i : j+1]
				}
			}
		}
	}
	return ans
}
