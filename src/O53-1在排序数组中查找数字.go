package main

func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i+j)/2
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	right := i
	i, j = 0, len(nums)-1
	for i <= j {
		mid := (i+j)/2
		if nums[mid] >= target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	left := j
	return right-left-1
}