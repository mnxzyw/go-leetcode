package main

import "fmt"

// 常规思路，遍历出所有无重复子串并记录长度，最后return最大长度

// 大佬思路
// 作者：frxiuxian
// 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/go-zhi-xing-yong-shi-0-ms-nei-cun-xiao-hao-27-mb-b/

func lengthOfLongestSubstring(s string) int {
	val := []byte(s)
	kvMap := make([]int, 128)
	lens := len(s)
	var max, num int
	// 分析
	// kvMap（其实是切片）的 key 为 字母的ascii码值，value 为 字母所在字符串的位置下标
	// i 表示可能出现的 重复元素 在字符串中的位置下标
	// j 表示遍历的游标
	for i, j := 0, 0; i < lens && j < lens; j++ {
		// 如果重复元素出现，后出现的元素位置下标一定大于先出现的
		if kvMap[val[j]] > i {
			// 更新重复元素位置，以便后续计算max长度
			i = kvMap[val[j]]
		}
		// 游标 减去 上个不重复元素的位置下标+1 = 距离（即长度）
		num = j - i + 1
		if num > max {
			max = num
		}
		// 将字母在字符串中的位置下标赋值到 相应字母在切片中的位置 所对应的 值（取巧点在于 切片的位置下标恰好对应相应字母的ascii码值）
		kvMap[val[j]] = j + 1
	}
	return max
}

func main() {
	s := "qweiqopqab"
	len := lengthOfLongestSubstring(s)
	fmt.Println(len)
}
