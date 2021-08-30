package main

import "fmt"

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	f1, f2 := 1, 2
	for i := 1; i < n; i++ {
		f1, f2 = f2, (f1+f2)%1000000007
	}
	return f1
}

func main() {
	fmt.Println(numWays(2))
	fmt.Println(numWays(7))
	fmt.Println(numWays(0))
}