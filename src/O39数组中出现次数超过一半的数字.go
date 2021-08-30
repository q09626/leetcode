package main

import "fmt"

func majorityElement(nums []int) int {
	x, vote := 0, 0
	for _, v := range nums {
		if vote == 0 {
			x = v
		}
		if v == x {
			vote++
		} else {
			vote--
		}
	}
	return x
}

func main() {
	fmt.Println(majorityElement([]int{1,2,3,2,2,2,5,4,2}))
}