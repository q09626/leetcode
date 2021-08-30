package main

import "fmt"

func findRepeatNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}

func main() {
	//nums := make([]int, 3, 5)
	nums := []int{1,2,3,3,0}
	fmt.Println(findRepeatNumber(nums))
}