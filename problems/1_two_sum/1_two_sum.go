package main

import "fmt"

func twoSum(nums []int, target int) []int {
	rsp := make([]int, 2)
	for key, _ := range nums {
		for i := key + 1; i < len(nums); i++ {
			if target == nums[key]+nums[i] {
				rsp[0] = key
				rsp[1] = i
				break
			}
		}
	}
	return rsp
}

func twoSum1(nums []int, target int) []int {
	tmp := make(map[int]int)
	rsp := make([]int, 2)
	for key, value := range nums {
		tmp[value] = key
	}
	for key, value := range nums {
		tofind := target - value
		index, ok := tmp[tofind]
		if ok && index != key {
			rsp[0] = key
			rsp[1] = index
			break
		}
	}
	return rsp
}

func main() {
	var nums = []int{1, 5, 9, 12, 17}
	var target = 17
	rsp := twoSum(nums, target)
	fmt.Println("rsp[0]", rsp[0], "rsp[1]", rsp[1])

	rsp1 := twoSum1(nums, target)
	fmt.Println("rsp1[0]", rsp1[0], "rsp1[1]", rsp1[1])
}
