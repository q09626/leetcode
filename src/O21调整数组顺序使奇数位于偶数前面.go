package main

func exchange(nums []int) []int {
	n := len(nums)
	low, high := 0, n-1
	for low < high {
		for low < n && nums[low] % 2 == 1 {
			low++
		}
		for high > -1 && nums[high] % 2 == 0 {
			high--
		}
		if low < high {
			nums[low], nums[high] = nums[high], nums[low]
		}
	}
	return nums
}

func main() {
	nums := []int{1,2,3,4}
	nums = exchange(nums)
	showNums(nums)
}