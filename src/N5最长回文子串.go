package main

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	ans := ""
	for l := 0; l <= len(s); l++ {
		for i := 0; i+l < len(s); i++ {
			j := i + l
			if s[i] == s[j] || (j-i <= 2 && dp[i+1][j-1]) {
				dp[i][j] = true
				continue
			}
			if l+1 > len(ans) {
				ans = s[i : l+1]
			}
		}
	}
	return ans
}
