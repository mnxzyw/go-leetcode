package main

import "fmt"

// 暴力遍历
// 1. 遍历出所有子串,
// 2. 判定每个子串是不是回文
func judgePalindrome(s string) bool {
	if len(s)%2 == 0 {
		var i, j = len(s)/2 - 1, len(s) / 2
		for ; i >= 0 && j <= len(s); i, j = i-1, j+1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	} else {
		var i, j = len(s) / 2, len(s) / 2
		for ; i >= 0 && j <= len(s); i, j = i-1, j+1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
}

func longestPalindrome(s string) string {
	var i, j int
	var smax []byte
	lmax := 0
	for ; i < len(s); i++ {
		for j = i + 1; j <= len(s); j++ {
			//[i,j]为左闭右开，注意循环范围取值
			str := s[i:j]
			if judgePalindrome(str) {
				if len(str) > lmax {
					lmax = len(str)
					smax = []byte(str)
				}
			}
		}
	}
	return string(smax)
}

// 大佬思路
// 动态规划
//func longestPalindrome1(s string) string {
//
//}

func main() {
	st := "a"
	sm := longestPalindrome(st)
	fmt.Println(sm)
}
