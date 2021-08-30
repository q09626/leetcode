package main

import "strconv"

func printNumbers(n int) []int {
	var num []byte
	var res []int
	nine, start := 0, n-1
	for i := 0; i < n; i++ {
		num = append(num, '0')
	}

	var dfs func(x int)
	dfs = func(x int) {
		if x == n {
			s := num[start:]
			if s[0] != '0' {
				sInt, _ := strconv.Atoi(string(s))
				res = append(res, sInt)
			}
			if n-start == nine {
				start -= 1
			}
			return
		}
		for i := 0; i < 10; i++ {
			if i == 9 {
				nine += 1
			}
			num[x] = strconv.Itoa(i)[0]
			dfs(x+1)
		}
		nine -= 1
	}

	dfs(0)
	return res
}

func main() {
	showNums(printNumbers(2))
}