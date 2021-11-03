package main

func permutation(s string) []string {
	n := len(s)
	ans := make([]string, 0) // 保存结果
	str := []byte(s)         // 原地修改str，省略中间结果

	var core func(str []byte, i int)
	core = func(str []byte, i int) {
		if i == n-1 {
			ans = append(ans, string(str))
			return
		}
		st := make(map[byte]int)
		for k := i; k < n; k++ {
			if _, ok := st[str[k]]; ok {
				continue
			}
			st[str[k]] = 1
			str[i], str[k] = str[k], str[i]
			core(str, i+1)
			str[i], str[k] = str[k], str[i]
		}
	}

	core(str, 0)
	return ans
}
