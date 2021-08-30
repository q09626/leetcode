package main

import "fmt"

func minArray(nums []int) int {
	low := 0
	high := len(nums)-1
	for low < high {
		mid := (high + low) / 2
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return nums[low]
}

func main() {
	fmt.Println(minArray([]int{3,4,5,1,2}))
	fmt.Println(minArray([]int{2,2,2,0,1}))
}