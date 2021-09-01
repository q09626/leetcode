package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	k := n / 2
	if n%2 == 1 {
		return float64(getKthElement(nums1, nums2, k+1)) // k是下标，k+1是第x个元素，传进去的是第X个元素
	} else {
		return float64(getKthElement(nums1, nums2, k)+getKthElement(nums1, nums2, k+1)) / 2.0
	}
}

func getKthElement(num1, num2 []int, k int) int {
	i, j := 0, 0
	for {
		if i == len(num1) {
			return num2[j+k-1] // 第x个元素，所以要减1
		}
		if j == len(num2) {
			return num1[i+k-1]
		}
		if k == 1 { // 第1个元素，所以减0
			return min(num1[i], num2[j])
		}
		m := k / 2                    // 不用管奇偶，就是取一半
		im := min(i+m, len(num1)) - 1 // 然后将这一半加到i和j上面
		jm := min(j+m, len(num2)) - 1 // 使得[i,im],[j,jm]的和等于k个元素
		e1, e2 := num1[im], num2[jm]
		if e1 <= e2 {
			k -= im - i + 1 // [i,im]的元素都不要了
			i = im + 1      // i从im的下一个元素开始
		} else {
			k -= jm - j + 1 // 上面同理
			j = jm + 1
		}
	}
}
