package main

func getLeastNumbers(nums []int, k int) []int {
	var core func(low int, high int)
	var partition func(low int, high int) int
	core = func(low int, high int) {
		if low >= high {
			return
		}
		p := partition(low, high)

		if p == k-1 {
			return
		} else if p > k {
			core(low, p-1)
		} else {
			core(p+1, high)
		}
	}
	partition = func(low int, high int) int {
		tmp := nums[low]
		for low < high {
			for low < high && nums[high] >= tmp {
				high--
			}
			nums[low] = nums[high]
			for low < high && nums[low] <= tmp {
				low++
			}
			nums[high] = nums[low]
		}
		nums[low] = tmp
		return low
	}

	core(0, len(nums)-1)
	return nums[:k]
}
