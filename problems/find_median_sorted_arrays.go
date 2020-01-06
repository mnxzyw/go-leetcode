package main

import "fmt"

// 归并排序的合并步骤
// 不知道大佬解法啥样 0.0
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lm := len(nums1) + len(nums2)
	s := make([]int, len(nums1)+len(nums2))
	var i, j, m int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			s[m] = nums1[i]
			i++
		} else {
			s[m] = nums2[j]
			j++
		}
		m++
	}
	for i < len(nums1) {
		s[m] = nums1[i]
		i++
		m++
	}
	for j < len(nums2) {
		s[m] = nums2[j]
		j++
		m++
	}
	if 1 == lm%2 {
		p := lm / 2
		return float64(s[p])
	} else {
		p := lm / 2
		return (float64(s[p-1]) + float64(s[p])) / 2
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	middle := findMedianSortedArrays(nums1, nums2)
	fmt.Println(middle)
}
