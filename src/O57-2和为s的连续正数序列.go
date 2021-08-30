package main

func findContinuousSequence(target int) [][]int {
	l, r := 1, 2
	var res [][]int
	for l < r && r < target {
		sum := (l+r)*(r-l+1)/2
		if sum == target {
			var tmp []int
			for i := l; i <= r; i++ {
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
			r++
			l++
		} else if sum > target {
			l++
		} else {
			r++
		}
	}
	return res
}

func main() {
	res := findContinuousSequence(15)
	for _, v := range res {
		showNums(v)
	}
}