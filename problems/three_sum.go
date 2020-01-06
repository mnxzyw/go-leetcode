package main

import "fmt"

func threeSum(nums []int) [][]int {
	li := len(nums)
	rsp := make([][]int, 0, li)

	for i := 0; i < li-2; i++ {
		for j := i + 1; j < li-1; j++ {
			for m := i + 2; m < li; m++ {
				if nums[i]+nums[j]+nums[m] == 0 {
					tmp := []int{nums[i], nums[j], nums[m]}
					rsp = append(rsp, tmp)
				}
			}
		}
	}
	return rsp
}

func main() {
	input := []int{-1, -2, -3, 0, 1, 2, 3}

	rsp := threeSum(input)
	fmt.Printf("%v", rsp)
}
