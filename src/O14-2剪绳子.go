package main

func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	p := 1000000007
	ans := 1
	for i := 0; i < a-1; i++ {
		ans = ans * 3 % p
	}
	if b == 0 {
		return int(ans * 3 % p)
	}
	if b == 1 {
		return int(ans * 4 % p)
	}
	return int(ans * 3 * 2 % p)
}

// 求(x^a)%p
func remainder(x, a, p int) int {
	ans := 1
	for i := 0; i < a; i++ {
		ans = (ans * x) % p
	}
	return ans
}
